from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

from ansible import constants as C
from ansible.playbook.task_include import TaskInclude
from ansible.plugins.callback import CallbackBase
from ansible.utils.color import colorize, hostcolor
import os

import requests

class CallbackModule(CallbackBase):

    '''
    '''

    CALLBACK_VERSION = 2.0
    CALLBACK_TYPE = 'notification'
    #  CALLBACK_TYPE = 'stdout'
    CALLBACK_NAME = 'pansible'


    def __init__(self):
        self._play = None
        self._last_task_banner = None
        super(CallbackModule, self).__init__()
        self.task_id = os.environ.get('PANSIBLE_RUN_ID')
        self.job_id = os.environ.get('PANSIBLE_JOB_ID')
        self.runner = os.environ.get('PANSIBLE_RUNNER')
        print(os.environ)

    def post(self, data):
        data['_pansible_task'] = self.task_id
        data['_pansible_job'] = self.job_id
        requests.post('http://%s/events' % self.runner, json=data)

    def v2_runner_on_failed(self, result, ignore_errors=False):
        self.post({
            "target": result._host.get_name(),
            "type": "failed",
            "result": result._result,
        })

    def v2_runner_on_ok(self, result):
        self.post({
            "target": result._host.get_name(),
            "type": "changed" if result._result.get('changed', False) else "ok",
            "result": result._result,
            "task": result._task.serialize(),
        })

    def v2_runner_on_skipped(self, result):
        self.post({
            "target": result._host.get_name(),
            "type": "skipped",
            "result": result._result,
        })

    def v2_runner_on_unreachable(self, result):
        self.post({
            "target": result._host.get_name(),
            "type": "unreachable",
            "result": result._result,
        })

    def v2_playbook_on_no_hosts_matched(self):
        self.post({
            "type": "no_hosts_matched",
        })

        self.post({
            "type": "finish",
            "reason": "no hosts matched",
        })


    def v2_playbook_on_no_hosts_remaining(self):
        self.post({
            "type": "no_hosts_remaining",
        })

        self.post({
            "type": "finish",
            "reason": "no hosts remaining",
        })

    def v2_playbook_on_task_start(self, task, is_conditional):
        self.post({
            "type": "task_start",
        })

    def v2_playbook_on_cleanup_task_start(self, task):
        self._display.banner("CLEANUP TASK [%s]" % task.get_name().strip())

    def v2_playbook_on_handler_task_start(self, task):
        self._display.banner("RUNNING HANDLER [%s]" % task.get_name().strip())

    def v2_playbook_on_play_start(self, play):
        self.post({
            "type": "play_start",
            "play": play.get_name(),
        })

    def v2_playbook_on_stats(self, stats):
        for h in sorted(stats.processed.keys()):
            t = stats.summarize(h)

            self.post({
                "target": h,
                "type": "stats",
                "ok": t["ok"],
                "changed": t["changed"],
                "unreachable": t["unreachable"],
                "failed": t["failures"],
            })

            self.post({
                "type": "finish",
                "reason": "success",
            })

    def v2_playbook_on_start(self, playbook):
        self.post({
            "type": "playbook_start",
            "file": playbook._file_name,
        })

    #  def v2_on_file_diff(self, result):
        #  if result._task.loop and 'results' in result._result:
            #  for res in result._result['results']:
                #  if 'diff' in res and res['diff'] and res.get('changed', False):
                    #  diff = self._get_diff(res['diff'])
                    #  if diff:
                        #  self._display.display(diff)
        #  elif 'diff' in result._result and result._result['diff'] and result._result.get('changed', False):
            #  diff = self._get_diff(result._result['diff'])
            #  if diff:
                #  self._display.display(diff)

