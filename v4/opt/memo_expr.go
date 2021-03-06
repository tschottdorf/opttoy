package opt

// memoExpr is a memoized representation of an expression. Specializations of
// memoExpr are generated by optgen for each operator (see Expr.og.go). Each
// memoExpr belongs to a memo group, which contain logically equivalent
// expressions. Two expressions are considered logically equivalent if they
// both reduce to an identical normal form after normalizing transformations
// have been applied.
//
// The children of memoExpr are recursively memoized in the same way as the
// memoExpr, and are referenced by their memo group. Therefore, the memoExpr
// is the root of a forest of expressions. Each memoExpr is memoized by its
// fingerprint, which is the hash of its op type plus the group ids of its
// children.
//
// Don't change the order of the fields in memoExpr. The op field is second in
// order to make the generated fingerprint methods faster and easier to
// implement.
type memoExpr struct {
	// group identifies the memo group to which this expression belongs.
	group GroupID

	// op is this expression's operator type. Each operator may have additional
	// fields. To access these fields, use the asXXX() generated methods to
	// cast the memoExpr to the more specialized expression type.
	op Operator
}
