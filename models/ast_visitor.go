package models

type IASTVisitor interface {
	// Statements
	VisitBlockNode(*ASTBlockNode) any
	VisitExprStmtNode(*ASTExprStmtNode) any
	VisitIfNode(*ASTIfNode) any
	VisitSwitchNode(*ASTSwitchNode) any
	VisitCaseNode(*ASTCaseNode) any
	VisitWhileNode(*ASTWhileNode) any
	VisitDoWhileNode(*ASTDoWhileNode) any
	VisitForNode(*ASTForNode) any
	VisitBreakNode(*ASTBreakNode) any
	VisitContinueNode(*ASTContinueNode) any
	VisitGotoNode(*ASTGotoNode) any
	VisitLabelNode(*ASTLabelNode) any
	VisitReturnNode(*ASTReturnNode) any

	// Expressions
	VisitAssignNode(*ASTAssignNode) any
	VisitOpAssignNode(*ASTOpAssignNode) any
	VisitCondExprNode(*ASTCondExprNode) any
	VisitLogicalOrNode(*ASTLogicalOrNode) any
	VisitLogicalAndNode(*ASTLogicalAndNode) any
	VisitBinaryOpNode(*ASTBinaryOpNode) any
	VisitUnaryOpNode(*ASTUnaryOpNode) any
	VisitPrefixOpNode(*ASTPrefixOpNode) any
	VisitSuffixOpNode(*ASTSuffixOpNode) any
	VisitArrayIdxRefNode(*ASTArrayIdxRefNode) any
	VisitMemberNode(*ASTMemberNode) any
	VisitPtrMemberNode(*ASTPtrMemberNode) any
	VisitFunctionCallNode(*ASTFunctionCallNode) any
	VisitDereferenceNode(*ASTDereferenceNode) any
	VisitAddressNode(*ASTAddressNode) any
	VisitCastNode(*ASTCastNode) any
	VisitSizeofExprNode(*ASTSizeofExprNode) any
	VisitSizeofTypeNode(*ASTSizeofTypeNode) any
	VisitVariableNode(*ASTVariableNode) any
	VisitIntegerLiteralNode(*ASTIntegerLiteralNode) any
	VisitStringLiteralNode(*ASTStringLiteralNode) any
}
