package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/hcl/token"
)

type Hcl struct {
	Resources []*ast.ObjectItem
}

type Resource struct {
	Item *ast.ObjectItem
}

func (h *Hcl) Encode(input map[string]interface{}) {

	resources, ok := input["resource"]
	if !ok {
		fmt.Fprintf(os.Stderr, "Hcl: Key resource not found in input data.")
		return
	}

	for name, resource := range resources.(map[string]interface{}) {
		for resourcekey, attributes := range resource.(map[string]interface{}) {
			res := NewResource(name, resourcekey)
			for attrname, attrvalue := range attributes.(map[string]interface{}) {
				res.AddAttribute(attrname, attrvalue.(string))
			}
			h.AddResource(res)
		}
	}

	root := &ast.File{
		Node: &ast.ObjectList{
			Items: h.Resources,
		}}
	printer.Fprint(os.Stdout, root)
}

func (h *Hcl) AddResource(r *Resource) {
	h.Resources = append(h.Resources, r.Item)
}

//AddResource build resource AST node
func NewResource(name, key string) *Resource {
	resource := &ast.ObjectItem{
		Keys: []*ast.ObjectKey{
			&ast.ObjectKey{
				Token: token.Token{Type: token.IDENT, Pos: token.Pos{Filename: "", Line: 1}, Text: "resource", JSON: false},
			},
			&ast.ObjectKey{
				Token: token.Token{Type: token.STRING, Pos: token.Pos{Filename: "", Line: 1}, Text: fmt.Sprintf("\"%s\"", name), JSON: false},
			},
			&ast.ObjectKey{
				Token: token.Token{Type: token.STRING, Pos: token.Pos{Filename: "", Line: 1}, Text: fmt.Sprintf("\"%s\"", key), JSON: false},
			},
		},
		Val: &ast.ObjectType{
			List: &ast.ObjectList{
				Items: []*ast.ObjectItem{},
			},
		},
	}
	return &Resource{Item: resource}
}

//AddResource add configuration argument to resource
func (r *Resource) AddAttribute(name, value string) *Resource {
	line := len(r.Item.Val.(*ast.ObjectType).List.Items) + 1
	r.Item.Val.(*ast.ObjectType).List.Items = append(r.Item.Val.(*ast.ObjectType).List.Items,
		&ast.ObjectItem{
			Keys: []*ast.ObjectKey{
				&ast.ObjectKey{Token: token.Token{Type: token.IDENT, Pos: token.Pos{Line: line}, Text: name, JSON: false}},
			},
			Assign: token.Pos{Filename: "", Line: line},
			Val: &ast.LiteralType{
				Token: token.Token{Type: token.STRING, Pos: token.Pos{Filename: "", Line: line}, Text: fmt.Sprintf("\"%s\"", value), JSON: false}},
		},
	)
	return r
}
