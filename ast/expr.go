package ast

type ExprVisitor interface {
	VisitBinaryExpr(BinaryExpr) LiteralType
	VisitUnaryExpr(UnaryExpr) LiteralType
	VisitToken(Token) LiteralType
}

type Expr interface {
	Accept(ExprVisitor) LiteralType
}

type BinaryExpr struct {
	Left  Expr
	Op    Token
	Right Expr
}

func (r BinaryExpr) Accept(v ExprVisitor) LiteralType {
	return v.VisitBinaryExpr(r)
}

type UnaryExpr struct {
	Op   Token
	Expr Expr
}

func (r UnaryExpr) Accept(v ExprVisitor) LiteralType {
	return v.VisitUnaryExpr(r)
}

func (r Token) Accept(v ExprVisitor) LiteralType {
	return v.VisitToken(r)
}
