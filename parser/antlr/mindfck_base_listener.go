// Code generated from parser/antlr/mindfck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // mindfck

import "github.com/antlr4-go/antlr/v4"

// BasemindfckListener is a complete listener for a parse tree produced by mindfckParser.
type BasemindfckListener struct{}

var _ mindfckListener = &BasemindfckListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasemindfckListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasemindfckListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasemindfckListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasemindfckListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatements is called when production statements is entered.
func (s *BasemindfckListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BasemindfckListener) ExitStatements(ctx *StatementsContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasemindfckListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasemindfckListener) ExitStatement(ctx *StatementContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasemindfckListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasemindfckListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasemindfckListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasemindfckListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasemindfckListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasemindfckListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrint is called when production print is entered.
func (s *BasemindfckListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BasemindfckListener) ExitPrint(ctx *PrintContext) {}

// EnterOperand is called when production operand is entered.
func (s *BasemindfckListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BasemindfckListener) ExitOperand(ctx *OperandContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasemindfckListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasemindfckListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasemindfckListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasemindfckListener) ExitLiteral(ctx *LiteralContext) {}
