package models

type IASTVisitor interface {
	VisitBlock(*ASTBlockNode) interface{}
	VisitBinaryOp(*ASTBinaryOpNode) interface{}
	VisitReturnNode(*ASTReturnNode) interface{}
	VisitExprStmtNode(*ASTExprStmtNode) interface{}
	VisitAssignNode(*ASTAssignNode) interface{}
	VisitOpAssignNode(*ASTOpAssignNode) interface{}
	VisitCondExprNode(*ASTCondExprNode) interface{}
	VisitLogicalOrNode(*ASTLogicalOrNode) interface{}
	VisitLogicalAndNode(*ASTLogicalAndNode) interface{}
	VisitUnaryOpNode(*ASTUnaryOpNode) interface{}
	VisitPrefixOpNode(*ASTPrefixOpNode) interface{}
	VisitDereferenceNode(*ASTDereferenceNode) interface{}
	VisitAddressNode(*ASTAddressNode) interface{}
	VisitSizeofTypeNode(*ASTSizeofTypeNode) interface{}
	VisitSizeofExprNode(*ASTSizeofExprNode) interface{}
	VisitSuffixOpNode(*ASTSuffixOpNode) interface{}
	VisitArrayIdxRefNode(*ASTArrayIdxRefNode) interface{}
	VisitFunctionCallNode(*ASTFunctionCallNode) interface{}
	VisitMemberNodeNode(*ASTMemberNode) interface{}
	VisitPtrMemberNode(*ASTPtrMemberNode) interface{}
	VisitIntegerLiteralNode(*ASTIntegerLiteralNode) interface{}
}
