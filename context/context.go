package context

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var Context = NewContext()

type RmContext struct {
	Request  *http.Request
	W        http.ResponseWriter
	routers  map[string]func(ctx *RmContext)
	pathArgs map[string]map[string]string
}

func NewContext() *RmContext {
	ctx := &RmContext{}
	ctx.routers = make(map[string]func(ctx2 *RmContext))
	ctx.pathArgs = make(map[string]map[string]string)
	return ctx
}

var UrlTree = NewTrie()

type Trie struct {
	next   map[string]*Trie
	isWord bool
}

func NewTrie() Trie {
	root := new(Trie)
	root.next = make(map[string]*Trie)
	root.isWord = false
	return *root
}

// 根据/拆分
func (t *Trie) Insert(word string) {
	for _, v := range strings.Split(word, "/") {
		if t.next[v] == nil {
			node := new(Trie)
			node.next = make(map[string]*Trie)
			node.isWord = false
			t.next[v] = node
		}

		if v == "*" || strings.Index(v, "{") != -1 {
			t.isWord = true
		}
		t = t.next[v]
	}
	t.isWord = true
}

func (t *Trie) Search(word string) (IsHave bool, args map[string]string) {
	args = make(map[string]string)
	IsHave = false
	for _, v := range strings.Split(word, "/") {
		if t.isWord {
			for k, _ := range t.next {
				if strings.Index(k, "{") != -1 {
					key := strings.Replace(k, "{", "", -1)
					key = strings.Replace(key, "}", "", -1)
					args[key] = v
				}
				v = k
			}
		}
		if t.next[v] == nil {
			log.Println("无法匹配路径")
			return
		}
		t = t.next[v]
	}
	if len(t.next) == 0 {
		IsHave = t.isWord
		return
	}
	return
}
func (ctx *RmContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx.W = w
	ctx.Request = r
	path := r.URL.Path
	f := ctx.routers[path]
	if f == nil {
		for key, value := range ctx.routers {
			//正则校验
			reg, _ := regexp.Compile("(/\\w+)*(/{\\w+})+(/\\w+)*")
			matchString := reg.MatchString(key)
			if !matchString {
				continue
			}
			isHave, args := UrlTree.Search(path)
			if isHave {
				//如果有存储在对应的路径上
				ctx.pathArgs[path] = args
				value(ctx)
			}
		}
	} else {
		f(ctx)
	}
}
func (ctx *RmContext) Handler(url string, f func(context *RmContext)) {
	//放入前缀树
	UrlTree.Insert(url)
	ctx.routers[url] = f
}

func (ctx *RmContext) GetPathVariable(key string) string {
	path := ctx.Request.URL.Path
	result := ctx.pathArgs[path][key]
	return result
}
func (ctx *RmContext) GetForm(key string) (string, error) {
	if err := ctx.Request.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		return "", err
	}
	return ctx.Request.Form.Get(key), nil
}
func (ctx *RmContext) GetJson(key string) interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	_ = json.Unmarshal(body, &params)
	return params[key]
}
