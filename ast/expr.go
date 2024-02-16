package ast

type ExprVisitor interface {
	VisitBinaryExp(BinaryExpr) Type
	VisitUnaryExp(UnaryExpr) Type
	VisitToken(Token) Type
}

type Expr interface {
	Accept(ExprVisitor) Type
}

type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (r BinaryExpr) String() string {
	return PrefixTraversal(r)
}

func (r BinaryExpr) Accept(v ExprVisitor) (typ Type) {
	typ = v.VisitBinaryExp(r)
	return
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (r UnaryExpr) Accept(v ExprVisitor) (typ Type) {
	typ = v.VisitUnaryExp(r)
	return
}

func (r UnaryExpr) String() string {
	return PrefixTraversal(r)
}

func (r Token) Accept(v ExprVisitor) (typ Type) {
	typ = v.VisitToken(r)
	return
}
