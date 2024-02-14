package ast

type ExprVisitor interface {
	VisitBinaryExp(BinaryExpr)
	VisitUnaryExp(UnaryExpr)
	VisitToken(Token)
}

type Expr interface {
	Accept(ExprVisitor)
}

type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (r BinaryExpr) Accept(visitor ExprVisitor) {
	visitor.VisitBinaryExp(r)
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (r UnaryExpr) Accept(visitor ExprVisitor) {
	visitor.VisitUnaryExp(r)
}

func (r Token) Accept(visitor ExprVisitor) {
	visitor.VisitToken(r)
}
