/*
Copyright © 2024 clarkmonkey@163.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/stkali/utility/errors"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stkali/utility/log"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "load filesystem folder or file to clix template",
	Run:   loadFunc,
	Args:  cobra.ExactArgs(1),
}

/*
load name@ 创建空的模板
load name@:README.md    		             -> ./README.md
load name@:                                  -> error
load name@:../Makefile						 -> Makefile
load name@:.								 -> 当前目录下整个导入模板
load mame@:src/								 -> ./src
load name@:src/main.py						 -> ./main.py
load name@:src/lib						     -> ./lib
load name@{{name}}.html:index.html			 -> .{{name}}.html
load name@src/{{name}}/{{name}}.py:main.py   -> ./src/{{name}}/{{name}}.py
load name@.:								 -> error
load name@.:.								 -> 将当前目录整个导入模板（带目录）
load name@.:src/main.py                      -> ./main.py
load name@.:src/lib							 -> ./lib/

load name@src/utils/index.py:../../index.py  -> ./src/urils/index.py
load name@src/:template/src/  		         -> ./src/

modify name@src/xxxxx

remove name@								 -> 删除模板 name
remove name@.								 -> 清空模板
remove name@src								 -> 删除src目录
remove name@src/main.py						 -> 删除模板中 src/main.py 文件

copy name@

create name:


table template:
id
name

*/

var (
	TemplateSep       = "@"
	MapSep            = ":"
	PermSep           = "!"
	PermMode          =0o644
	InvalidTokenError = errors.New("invalid token")
)

// loadFunc ...
// 解析load参数，解析token，调用不同的实现
func loadFunc(cmd *cobra.Command, args []string) {
	log.Infof("subcommand load, os.Args: %s, subcommand args: %s", os.Args, args)
	_, _ = checkToken(args[0])
	//
	//templateName := token[:tempSepIndex]
	//if tempSepIndex+1 == len(token) {
	//	err := internal.CreateTemplate(templateName)
	//	errors.CheckErr(err)
	//	log.Infof("Successfully created template: %s", templateName)
	//	return
	//}
	//
	//// 有映射关系的
	//mapping := token[tempSepIndex+1:]
	//log.Infof("mapping string: %s", mapping)
	//mapSlice := strings.Split(mapping, MapSep)
	//if len(mapSlice) != 2 {
	//	log.Errorf("invalid mapping: %s", mapping)
	//	return
	//}
	//src := mapSlice[0]
	//dst := mapSlice[1]
	//// 缺省了项目名
	//log.Infof("src: %s, dst: %s", src, dst)
	//matches, err := filepath.Glob(dst)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//if len(matches) == 0 {
	//	log.Info("empty file list")
	//	return
	//}

}

// checkToken ...
// name@
// name@:
func checkToken(token string) (string, error) {
	tempSepIndex := strings.Index(token, TemplateSep)
	// not found @
	if tempSepIndex == -1 {
		log.Infof("invalid token: %q not found %s", token, TemplateSep)
		return "", InvalidTokenError
	}
	mapping := token[tempSepIndex:]
	tmp := strings.Split(mapping, MapSep)
	// check src and dst
	if len(tmp) != 2 {
		log.Infof("invalid mapping: %q, contains %d map sep", mapping, strings.Contains(mapping, MapSep))
		return "", InvalidTokenError
	}
	src, dst := tmp[0], tmp[1]
	// empty dst
	if dst == "" {
		log.Infof("invalid mapping: %q, dst path is empty", mapping)
		return "", InvalidTokenError
	}
	// invalid permission string
	fmt.Println(src)
	return dst, nil

}

type Path struct {
	P string
	Perm int
}

type LoadMapping struct {
	Src Path
	Dst Path
}

func NewLoadMapping(mapping string) (*LoadMapping, error) {
	tmp := strings.Split(mapping, MapSep)
	if len(tmp) != 2 {
		return nil, InvalidTokenError
	}
	// 
	srcPath, dstPath := tmp[0], tmp[1]
	if srcPath == "" {
		return nil, InvalidTokenError
	}
	log.Infof("dst path: %s, src path: %s", dstPath, srcPath)
	lmap := &LoadMapping{}
	
	return lmap, nil
}



func parseLoadToken(token string) (*LoadMapping, error)  {

	tempSepIndex := strings.Index(token, TemplateSep)
	// not found @
	if tempSepIndex == -1 {
		log.Infof("invalid token: %q not found %s", token, TemplateSep)
		return nil, InvalidTokenError
	}
	return NewLoadMapping(token[tempSepIndex:])
}

func init() {
	rootCmd.AddCommand(loadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func main() {
	fmt.Println("hello world!")
}
