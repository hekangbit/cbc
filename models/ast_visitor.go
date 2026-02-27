package models

// TODO：
type IASTVisitor interface {
	// Statements
	VisitBlock(*ASTBlockNode) interface{}
	VisitExprStmtNode(*ASTExprStmtNode) interface{}
	// VisitIfStmtNode
	// VisitSwitchNode
	// VisitCaseNode
	// VisitWhileNode
	// VisitDoWhileNode
	// VisitForNode
	// VisitBreakNode
	// VisitContinueNode
	// VisitGotoNode
	// VisitLabelNode
	VisitReturnNode(*ASTReturnNode) interface{}

	// Expressions
	VisitAssignNode(*ASTAssignNode) interface{}
	VisitOpAssignNode(*ASTOpAssignNode) interface{}
	VisitCondExprNode(*ASTCondExprNode) interface{}
	VisitLogicalOrNode(*ASTLogicalOrNode) interface{}
	VisitLogicalAndNode(*ASTLogicalAndNode) interface{}
	VisitBinaryOp(*ASTBinaryOpNode) interface{}
	VisitUnaryOpNode(*ASTUnaryOpNode) interface{}
	VisitPrefixOpNode(*ASTPrefixOpNode) interface{}
	VisitSuffixOpNode(*ASTSuffixOpNode) interface{}
	VisitArrayIdxRefNode(*ASTArrayIdxRefNode) interface{}
	VisitMemberNodeNode(*ASTMemberNode) interface{}
	VisitPtrMemberNode(*ASTPtrMemberNode) interface{}
	VisitFunctionCallNode(*ASTFunctionCallNode) interface{}
	VisitDereferenceNode(*ASTDereferenceNode) interface{}
	VisitAddressNode(*ASTAddressNode) interface{}
	// VisitCastNode
	VisitSizeofExprNode(*ASTSizeofExprNode) interface{}
	VisitSizeofTypeNode(*ASTSizeofTypeNode) interface{}
	VisitVariableNode(*ASTVariableNode) interface{}
	VisitIntegerLiteralNode(*ASTIntegerLiteralNode) interface{}
	VisitStringLiteralNode(*ASTStringLiteralNode) interface{}
}
