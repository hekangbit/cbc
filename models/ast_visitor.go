package models

// TODO：
type IASTVisitor interface {
	// Statements
	VisitBlockNode(*ASTBlockNode) any
	VisitExprStmtNode(*ASTExprStmtNode) any
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
	VisitReturnNode(*ASTReturnNode) any

	// Expressions
	VisitAssignNode(*ASTAssignNode) any
	VisitOpAssignNode(*ASTOpAssignNode) any
	VisitCondExprNode(*ASTCondExprNode) any
	VisitLogicalOrNode(*ASTLogicalOrNode) any
	VisitLogicalAndNode(*ASTLogicalAndNode) any
	VisitBinaryOp(*ASTBinaryOpNode) any
	VisitUnaryOpNode(*ASTUnaryOpNode) any
	VisitPrefixOpNode(*ASTPrefixOpNode) any
	VisitSuffixOpNode(*ASTSuffixOpNode) any
	VisitArrayIdxRefNode(*ASTArrayIdxRefNode) any
	VisitMemberNodeNode(*ASTMemberNode) any
	VisitPtrMemberNode(*ASTPtrMemberNode) any
	VisitFunctionCallNode(*ASTFunctionCallNode) any
	VisitDereferenceNode(*ASTDereferenceNode) any
	VisitAddressNode(*ASTAddressNode) any
	// VisitCastNode
	VisitSizeofExprNode(*ASTSizeofExprNode) any
	VisitSizeofTypeNode(*ASTSizeofTypeNode) any
	VisitVariableNode(*ASTVariableNode) any
	VisitIntegerLiteralNode(*ASTIntegerLiteralNode) any
	VisitStringLiteralNode(*ASTStringLiteralNode) any
}
