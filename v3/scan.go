package v3

import (
	"bytes"
)

func init() {
	registerOperator(scanOp, "scan", scan{})
}

func newScanExpr(tab *table) *expr {
	return &expr{
		op:      scanOp,
		private: tab,
	}
}

type scan struct{}

func (scan) kind() operatorKind {
	return logicalKind | relationalKind
}

func (scan) layout() exprLayout {
	return exprLayout{}
}

func (scan) format(e *expr, buf *bytes.Buffer, level int) {
	formatRelational(e, buf, level)
}

func (scan) initKeys(e *expr, state *queryState) {
	tab := e.private.(*table)
	props := e.props
	props.foreignKeys = nil

	for _, k := range tab.keys {
		if k.fkey != nil {
			base, ok := state.tables[k.fkey.referenced.name]
			if !ok {
				// The referenced table is not part of the query.
				continue
			}

			var src bitmap
			for _, i := range k.columns {
				src.Add(props.columns[i].index)
			}
			var dest bitmap
			for _, i := range k.fkey.columns {
				dest.Add(base + bitmapIndex(i))
			}

			props.foreignKeys = append(props.foreignKeys, foreignKeyProps{
				src:  src,
				dest: dest,
			})
		}
	}
}

func (scan) updateProps(e *expr) {
	e.props.outerCols = e.requiredInputCols().Difference(e.props.outputCols)
}

func (scan) requiredProps(required *physicalProps, child int) *physicalProps {
	return nil
}
