package version

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/frainzy1477/trojan-go/constant"
	"github.com/frainzy1477/trojan-go/option"

	"github.com/frainzy1477/trojan-go/common"
)

type versionOption struct {
	flag *bool
}

func (*versionOption) Name() string {
	return "version"
}

func (*versionOption) Priority() int {
	return 10
}

func (c *versionOption) Handle() error {
	if *c.flag {
		fmt.Println("Trojan-Go", constant.Version)
		fmt.Println("Go Version:", runtime.Version())
		fmt.Println("OS/Arch:", runtime.GOOS+"/"+runtime.GOARCH)
		fmt.Println("Git Commit:", constant.Commit)
		fmt.Println("Developed by PageFault")
		fmt.Println("Licensed under GNU General Public License version 3")
		return nil
	}
	return common.NewError("not set")
}

func init() {
	option.RegisterHandler(&versionOption{
		flag: flag.Bool("version", false, "Display version and help info"),
	})
}
