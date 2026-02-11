pub enum Operator {
    Plus,
    Minus,
}

pub enum Expr {
    Int(i32),
    UnaryExpr {
        op: Operator,
        child: Box<Expr>,
    },
    BinaryExpr {
        op: Operator,
        lhs: Box<Expr>,
        rhs: Box<Expr>,
    },
}
