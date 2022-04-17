package templates

import (
	"reflect"
	"text/template/parse"
)

type WalkF func(newNode parse.Node) (bool, error)
type Visitors struct {
	ActionNode     func(*parse.ActionNode) (bool, error)
	BoolNode       func(*parse.BoolNode) (bool, error)
	BranchNode     func(*parse.BranchNode) (bool, error)
	BreakNode      func(*parse.BreakNode) (bool, error)
	ChainNode      func(*parse.ChainNode) (bool, error)
	CommandNode    func(*parse.CommandNode) (bool, error)
	CommentNode    func(*parse.CommentNode) (bool, error)
	ContinueNode   func(*parse.ContinueNode) (bool, error)
	DotNode        func(*parse.DotNode) (bool, error)
	FieldNode      func(*parse.FieldNode) (bool, error)
	IdentifierNode func(*parse.IdentifierNode) (bool, error)
	IfNode         func(*parse.IfNode) (bool, error)
	ListNode       func(*parse.ListNode) (bool, error)
	NilNode        func(*parse.NilNode) (bool, error)
	NumberNode     func(*parse.NumberNode) (bool, error)
	PipeNode       func(*parse.PipeNode) (bool, error)
	RangeNode      func(*parse.RangeNode) (bool, error)
	StringNode     func(*parse.StringNode) (bool, error)
	TemplateNode   func(*parse.TemplateNode) (bool, error)
	TextNode       func(*parse.TextNode) (bool, error)
	VariableNode   func(*parse.VariableNode) (bool, error)
}

func Walk(n parse.Node, visitors Visitors) error {
	if reflect.ValueOf(n).IsNil() {
		return nil
	}

	cont, err := dispatch(n, visitors)

	if err != nil {
		return err
	}

	if !cont {
		return nil
	}

	next := []parse.Node{}

	switch n := n.(type) {
	case *parse.ActionNode:
		next = append(next, n.Pipe)
	case *parse.BranchNode:
		next = append(next, n.Pipe)
		next = append(next, n.List)
		next = append(next, n.ElseList)
	case *parse.CommandNode:
		next = append(next, n.Args...)
	case *parse.IfNode:
		next = append(next, &n.BranchNode)
	case *parse.ListNode:
		next = append(next, n.Nodes...)
	case *parse.PipeNode:
		for _, el := range n.Decl {
			next = append(next, el)
		}
		for _, el := range n.Cmds {
			next = append(next, el)
		}
	case *parse.RangeNode:
		next = append(next, &n.BranchNode)
	case *parse.TemplateNode:
		next = append(next, n.Pipe)
		// TODO: parse this template too?
	case *parse.WithNode:
		next = append(next, &n.BranchNode)
	}

	for _, el := range next {
		err := Walk(el, visitors)
		if err != nil {
			return err
		}
	}

	return err
}

func dispatch(n parse.Node, visitors Visitors) (cont bool, err error) {
	cont = true
	switch n := n.(type) {
	case *parse.ActionNode:
		if visitors.ActionNode != nil {
			cont, err = visitors.ActionNode(n)
		}
	case *parse.BoolNode:
		if visitors.BoolNode != nil {
			cont, err = visitors.BoolNode(n)
		}
	case *parse.BranchNode:
		if visitors.BranchNode != nil {
			cont, err = visitors.BranchNode(n)
		}
	case *parse.BreakNode:
		if visitors.BreakNode != nil {
			cont, err = visitors.BreakNode(n)
		}
	case *parse.ChainNode:
		if visitors.ChainNode != nil {
			cont, err = visitors.ChainNode(n)
		}
	case *parse.CommandNode:
		if visitors.CommandNode != nil {
			cont, err = visitors.CommandNode(n)
		}
	case *parse.CommentNode:
		if visitors.CommentNode != nil {
			cont, err = visitors.CommentNode(n)
		}
	case *parse.ContinueNode:
		if visitors.ContinueNode != nil {
			cont, err = visitors.ContinueNode(n)
		}
	case *parse.DotNode:
		if visitors.DotNode != nil {
			cont, err = visitors.DotNode(n)
		}
	case *parse.FieldNode:
		if visitors.FieldNode != nil {
			cont, err = visitors.FieldNode(n)
		}
	case *parse.IdentifierNode:
		if visitors.IdentifierNode != nil {
			cont, err = visitors.IdentifierNode(n)
		}
	case *parse.IfNode:
		if visitors.IfNode != nil {
			cont, err = visitors.IfNode(n)
		}
	case *parse.ListNode:
		if visitors.ListNode != nil {
			cont, err = visitors.ListNode(n)
		}
	case *parse.NilNode:
		if visitors.NilNode != nil {
			cont, err = visitors.NilNode(n)
		}
	case *parse.NumberNode:
		if visitors.NumberNode != nil {
			cont, err = visitors.NumberNode(n)
		}
	case *parse.PipeNode:
		if visitors.PipeNode != nil {
			cont, err = visitors.PipeNode(n)
		}
	case *parse.RangeNode:
		if visitors.RangeNode != nil {
			cont, err = visitors.RangeNode(n)
		}
	case *parse.StringNode:
		if visitors.StringNode != nil {
			cont, err = visitors.StringNode(n)
		}
	case *parse.TemplateNode:
		if visitors.TemplateNode != nil {
			cont, err = visitors.TemplateNode(n)
		}
	case *parse.TextNode:
		if visitors.TextNode != nil {
			cont, err = visitors.TextNode(n)
		}
	case *parse.VariableNode:
		if visitors.VariableNode != nil {
			cont, err = visitors.VariableNode(n)
		}
	}
	return
}
