package cmd

import (
	"fmt"
	"testing"

	"git.elenet.me/yuelong.huang/pansible/utils"
)

func TestParse(t *testing.T) {
	inv := runInventory{
		Hosts: utils.Map{
			"all": utils.Map{
				"hosts": []string{"mail.example.com"},
				"children": utils.Map{
					"webservers": utils.Map{
						"hosts": []string{
							"foo.ex",
							"bar",
						},
					},
					"dbserver": utils.Map{
						"hosts": []string{"abc.123.com"},
					},
				},
			},
		},
		Vars: utils.Map{
			"all": utils.Map{
				"a": 1,
			},
			"webservers": utils.Map{
				"b": 2,
			},
		},
	}

	res, err := getGroup(inv.Hosts)
	fmt.Println(res, err)
	setVars(res, inv.Vars)

	printJSON(res)
}
