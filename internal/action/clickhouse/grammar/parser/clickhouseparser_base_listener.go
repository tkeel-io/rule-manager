// Generated from ClickHouseParser.g4 by ANTLR 4.7.

package parser // ClickHouseParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseClickHouseParserListener is a complete listener for a parse tree produced by ClickHouseParser.
type BaseClickHouseParserListener struct{}

var _ ClickHouseParserListener = &BaseClickHouseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClickHouseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClickHouseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClickHouseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClickHouseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseClickHouseParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseClickHouseParserListener) ExitParse(ctx *ParseContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseClickHouseParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseClickHouseParserListener) ExitQuery(ctx *QueryContext) {}

// EnterSelect_query is called when production select_query is entered.
func (s *BaseClickHouseParserListener) EnterSelect_query(ctx *Select_queryContext) {}

// ExitSelect_query is called when production select_query is exited.
func (s *BaseClickHouseParserListener) ExitSelect_query(ctx *Select_queryContext) {}

// EnterSelect_query_main is called when production select_query_main is entered.
func (s *BaseClickHouseParserListener) EnterSelect_query_main(ctx *Select_query_mainContext) {}

// ExitSelect_query_main is called when production select_query_main is exited.
func (s *BaseClickHouseParserListener) ExitSelect_query_main(ctx *Select_query_mainContext) {}

// EnterSelect_with_step is called when production select_with_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_with_step(ctx *Select_with_stepContext) {}

// ExitSelect_with_step is called when production select_with_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_with_step(ctx *Select_with_stepContext) {}

// EnterSelect_select_step is called when production select_select_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_select_step(ctx *Select_select_stepContext) {}

// ExitSelect_select_step is called when production select_select_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_select_step(ctx *Select_select_stepContext) {}

// EnterSelect_from_step is called when production select_from_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_from_step(ctx *Select_from_stepContext) {}

// ExitSelect_from_step is called when production select_from_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_from_step(ctx *Select_from_stepContext) {}

// EnterSelect_array_join_step is called when production select_array_join_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_array_join_step(ctx *Select_array_join_stepContext) {
}

// ExitSelect_array_join_step is called when production select_array_join_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_array_join_step(ctx *Select_array_join_stepContext) {
}

// EnterSelect_sample_step is called when production select_sample_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_sample_step(ctx *Select_sample_stepContext) {}

// ExitSelect_sample_step is called when production select_sample_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_sample_step(ctx *Select_sample_stepContext) {}

// EnterSample_ratio is called when production sample_ratio is entered.
func (s *BaseClickHouseParserListener) EnterSample_ratio(ctx *Sample_ratioContext) {}

// ExitSample_ratio is called when production sample_ratio is exited.
func (s *BaseClickHouseParserListener) ExitSample_ratio(ctx *Sample_ratioContext) {}

// EnterSelect_join_step is called when production select_join_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_join_step(ctx *Select_join_stepContext) {}

// ExitSelect_join_step is called when production select_join_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_join_step(ctx *Select_join_stepContext) {}

// EnterSelect_join_right_part is called when production select_join_right_part is entered.
func (s *BaseClickHouseParserListener) EnterSelect_join_right_part(ctx *Select_join_right_partContext) {
}

// ExitSelect_join_right_part is called when production select_join_right_part is exited.
func (s *BaseClickHouseParserListener) ExitSelect_join_right_part(ctx *Select_join_right_partContext) {
}

// EnterSelect_prewhere_step is called when production select_prewhere_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_prewhere_step(ctx *Select_prewhere_stepContext) {}

// ExitSelect_prewhere_step is called when production select_prewhere_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_prewhere_step(ctx *Select_prewhere_stepContext) {}

// EnterSelect_where_step is called when production select_where_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_where_step(ctx *Select_where_stepContext) {}

// ExitSelect_where_step is called when production select_where_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_where_step(ctx *Select_where_stepContext) {}

// EnterSelect_groupby_step is called when production select_groupby_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_groupby_step(ctx *Select_groupby_stepContext) {}

// ExitSelect_groupby_step is called when production select_groupby_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_groupby_step(ctx *Select_groupby_stepContext) {}

// EnterSelect_having_step is called when production select_having_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_having_step(ctx *Select_having_stepContext) {}

// ExitSelect_having_step is called when production select_having_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_having_step(ctx *Select_having_stepContext) {}

// EnterSelect_orderby_step is called when production select_orderby_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_orderby_step(ctx *Select_orderby_stepContext) {}

// ExitSelect_orderby_step is called when production select_orderby_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_orderby_step(ctx *Select_orderby_stepContext) {}

// EnterSelect_limit_step is called when production select_limit_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_limit_step(ctx *Select_limit_stepContext) {}

// ExitSelect_limit_step is called when production select_limit_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_limit_step(ctx *Select_limit_stepContext) {}

// EnterSelect_limitby_step is called when production select_limitby_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_limitby_step(ctx *Select_limitby_stepContext) {}

// ExitSelect_limitby_step is called when production select_limitby_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_limitby_step(ctx *Select_limitby_stepContext) {}

// EnterSettings_step is called when production settings_step is entered.
func (s *BaseClickHouseParserListener) EnterSettings_step(ctx *Settings_stepContext) {}

// ExitSettings_step is called when production settings_step is exited.
func (s *BaseClickHouseParserListener) ExitSettings_step(ctx *Settings_stepContext) {}

// EnterSelect_format_step is called when production select_format_step is entered.
func (s *BaseClickHouseParserListener) EnterSelect_format_step(ctx *Select_format_stepContext) {}

// ExitSelect_format_step is called when production select_format_step is exited.
func (s *BaseClickHouseParserListener) ExitSelect_format_step(ctx *Select_format_stepContext) {}

// EnterInsert_query is called when production insert_query is entered.
func (s *BaseClickHouseParserListener) EnterInsert_query(ctx *Insert_queryContext) {}

// ExitInsert_query is called when production insert_query is exited.
func (s *BaseClickHouseParserListener) ExitInsert_query(ctx *Insert_queryContext) {}

// EnterCreate_query is called when production create_query is entered.
func (s *BaseClickHouseParserListener) EnterCreate_query(ctx *Create_queryContext) {}

// ExitCreate_query is called when production create_query is exited.
func (s *BaseClickHouseParserListener) ExitCreate_query(ctx *Create_queryContext) {}

// EnterRename_query is called when production rename_query is entered.
func (s *BaseClickHouseParserListener) EnterRename_query(ctx *Rename_queryContext) {}

// ExitRename_query is called when production rename_query is exited.
func (s *BaseClickHouseParserListener) ExitRename_query(ctx *Rename_queryContext) {}

// EnterDrop_query is called when production drop_query is entered.
func (s *BaseClickHouseParserListener) EnterDrop_query(ctx *Drop_queryContext) {}

// ExitDrop_query is called when production drop_query is exited.
func (s *BaseClickHouseParserListener) ExitDrop_query(ctx *Drop_queryContext) {}

// EnterAlter_query is called when production alter_query is entered.
func (s *BaseClickHouseParserListener) EnterAlter_query(ctx *Alter_queryContext) {}

// ExitAlter_query is called when production alter_query is exited.
func (s *BaseClickHouseParserListener) ExitAlter_query(ctx *Alter_queryContext) {}

// EnterAlter_query_element is called when production alter_query_element is entered.
func (s *BaseClickHouseParserListener) EnterAlter_query_element(ctx *Alter_query_elementContext) {}

// ExitAlter_query_element is called when production alter_query_element is exited.
func (s *BaseClickHouseParserListener) ExitAlter_query_element(ctx *Alter_query_elementContext) {}

// EnterClickhouse_type is called when production clickhouse_type is entered.
func (s *BaseClickHouseParserListener) EnterClickhouse_type(ctx *Clickhouse_typeContext) {}

// ExitClickhouse_type is called when production clickhouse_type is exited.
func (s *BaseClickHouseParserListener) ExitClickhouse_type(ctx *Clickhouse_typeContext) {}

// EnterSimple_type is called when production simple_type is entered.
func (s *BaseClickHouseParserListener) EnterSimple_type(ctx *Simple_typeContext) {}

// ExitSimple_type is called when production simple_type is exited.
func (s *BaseClickHouseParserListener) ExitSimple_type(ctx *Simple_typeContext) {}

// EnterEnum_entry is called when production enum_entry is entered.
func (s *BaseClickHouseParserListener) EnterEnum_entry(ctx *Enum_entryContext) {}

// ExitEnum_entry is called when production enum_entry is exited.
func (s *BaseClickHouseParserListener) ExitEnum_entry(ctx *Enum_entryContext) {}

// EnterUse_query is called when production use_query is entered.
func (s *BaseClickHouseParserListener) EnterUse_query(ctx *Use_queryContext) {}

// ExitUse_query is called when production use_query is exited.
func (s *BaseClickHouseParserListener) ExitUse_query(ctx *Use_queryContext) {}

// EnterSet_query is called when production set_query is entered.
func (s *BaseClickHouseParserListener) EnterSet_query(ctx *Set_queryContext) {}

// ExitSet_query is called when production set_query is exited.
func (s *BaseClickHouseParserListener) ExitSet_query(ctx *Set_queryContext) {}

// EnterAssignment_list is called when production assignment_list is entered.
func (s *BaseClickHouseParserListener) EnterAssignment_list(ctx *Assignment_listContext) {}

// ExitAssignment_list is called when production assignment_list is exited.
func (s *BaseClickHouseParserListener) ExitAssignment_list(ctx *Assignment_listContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseClickHouseParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseClickHouseParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterKill_query_query is called when production kill_query_query is entered.
func (s *BaseClickHouseParserListener) EnterKill_query_query(ctx *Kill_query_queryContext) {}

// ExitKill_query_query is called when production kill_query_query is exited.
func (s *BaseClickHouseParserListener) ExitKill_query_query(ctx *Kill_query_queryContext) {}

// EnterOptimize_query is called when production optimize_query is entered.
func (s *BaseClickHouseParserListener) EnterOptimize_query(ctx *Optimize_queryContext) {}

// ExitOptimize_query is called when production optimize_query is exited.
func (s *BaseClickHouseParserListener) ExitOptimize_query(ctx *Optimize_queryContext) {}

// EnterTable_properties_query is called when production table_properties_query is entered.
func (s *BaseClickHouseParserListener) EnterTable_properties_query(ctx *Table_properties_queryContext) {
}

// ExitTable_properties_query is called when production table_properties_query is exited.
func (s *BaseClickHouseParserListener) ExitTable_properties_query(ctx *Table_properties_queryContext) {
}

// EnterShow_tables_query is called when production show_tables_query is entered.
func (s *BaseClickHouseParserListener) EnterShow_tables_query(ctx *Show_tables_queryContext) {}

// ExitShow_tables_query is called when production show_tables_query is exited.
func (s *BaseClickHouseParserListener) ExitShow_tables_query(ctx *Show_tables_queryContext) {}

// EnterShow_processlist_query is called when production show_processlist_query is entered.
func (s *BaseClickHouseParserListener) EnterShow_processlist_query(ctx *Show_processlist_queryContext) {
}

// ExitShow_processlist_query is called when production show_processlist_query is exited.
func (s *BaseClickHouseParserListener) ExitShow_processlist_query(ctx *Show_processlist_queryContext) {
}

// EnterCheck_query is called when production check_query is entered.
func (s *BaseClickHouseParserListener) EnterCheck_query(ctx *Check_queryContext) {}

// ExitCheck_query is called when production check_query is exited.
func (s *BaseClickHouseParserListener) ExitCheck_query(ctx *Check_queryContext) {}

// EnterFull_table_name is called when production full_table_name is entered.
func (s *BaseClickHouseParserListener) EnterFull_table_name(ctx *Full_table_nameContext) {}

// ExitFull_table_name is called when production full_table_name is exited.
func (s *BaseClickHouseParserListener) ExitFull_table_name(ctx *Full_table_nameContext) {}

// EnterPartition_name is called when production partition_name is entered.
func (s *BaseClickHouseParserListener) EnterPartition_name(ctx *Partition_nameContext) {}

// ExitPartition_name is called when production partition_name is exited.
func (s *BaseClickHouseParserListener) ExitPartition_name(ctx *Partition_nameContext) {}

// EnterCluster_name is called when production cluster_name is entered.
func (s *BaseClickHouseParserListener) EnterCluster_name(ctx *Cluster_nameContext) {}

// ExitCluster_name is called when production cluster_name is exited.
func (s *BaseClickHouseParserListener) ExitCluster_name(ctx *Cluster_nameContext) {}

// EnterDatabase_name is called when production database_name is entered.
func (s *BaseClickHouseParserListener) EnterDatabase_name(ctx *Database_nameContext) {}

// ExitDatabase_name is called when production database_name is exited.
func (s *BaseClickHouseParserListener) ExitDatabase_name(ctx *Database_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseClickHouseParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseClickHouseParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterFormat_name is called when production format_name is entered.
func (s *BaseClickHouseParserListener) EnterFormat_name(ctx *Format_nameContext) {}

// ExitFormat_name is called when production format_name is exited.
func (s *BaseClickHouseParserListener) ExitFormat_name(ctx *Format_nameContext) {}

// EnterQuery_outfile_step is called when production query_outfile_step is entered.
func (s *BaseClickHouseParserListener) EnterQuery_outfile_step(ctx *Query_outfile_stepContext) {}

// ExitQuery_outfile_step is called when production query_outfile_step is exited.
func (s *BaseClickHouseParserListener) ExitQuery_outfile_step(ctx *Query_outfile_stepContext) {}

// EnterEngine is called when production engine is entered.
func (s *BaseClickHouseParserListener) EnterEngine(ctx *EngineContext) {}

// ExitEngine is called when production engine is exited.
func (s *BaseClickHouseParserListener) ExitEngine(ctx *EngineContext) {}

// EnterIdentifier_with_optional_parameters is called when production identifier_with_optional_parameters is entered.
func (s *BaseClickHouseParserListener) EnterIdentifier_with_optional_parameters(ctx *Identifier_with_optional_parametersContext) {
}

// ExitIdentifier_with_optional_parameters is called when production identifier_with_optional_parameters is exited.
func (s *BaseClickHouseParserListener) ExitIdentifier_with_optional_parameters(ctx *Identifier_with_optional_parametersContext) {
}

// EnterIdentifier_with_parameters is called when production identifier_with_parameters is entered.
func (s *BaseClickHouseParserListener) EnterIdentifier_with_parameters(ctx *Identifier_with_parametersContext) {
}

// ExitIdentifier_with_parameters is called when production identifier_with_parameters is exited.
func (s *BaseClickHouseParserListener) ExitIdentifier_with_parameters(ctx *Identifier_with_parametersContext) {
}

// EnterOrder_by_expression_list is called when production order_by_expression_list is entered.
func (s *BaseClickHouseParserListener) EnterOrder_by_expression_list(ctx *Order_by_expression_listContext) {
}

// ExitOrder_by_expression_list is called when production order_by_expression_list is exited.
func (s *BaseClickHouseParserListener) ExitOrder_by_expression_list(ctx *Order_by_expression_listContext) {
}

// EnterOrder_by_element is called when production order_by_element is entered.
func (s *BaseClickHouseParserListener) EnterOrder_by_element(ctx *Order_by_elementContext) {}

// ExitOrder_by_element is called when production order_by_element is exited.
func (s *BaseClickHouseParserListener) ExitOrder_by_element(ctx *Order_by_elementContext) {}

// EnterTable_ttl_list is called when production table_ttl_list is entered.
func (s *BaseClickHouseParserListener) EnterTable_ttl_list(ctx *Table_ttl_listContext) {}

// ExitTable_ttl_list is called when production table_ttl_list is exited.
func (s *BaseClickHouseParserListener) ExitTable_ttl_list(ctx *Table_ttl_listContext) {}

// EnterPartition_by_element is called when production partition_by_element is entered.
func (s *BaseClickHouseParserListener) EnterPartition_by_element(ctx *Partition_by_elementContext) {}

// ExitPartition_by_element is called when production partition_by_element is exited.
func (s *BaseClickHouseParserListener) ExitPartition_by_element(ctx *Partition_by_elementContext) {}

// EnterTable_ttl_declaration is called when production table_ttl_declaration is entered.
func (s *BaseClickHouseParserListener) EnterTable_ttl_declaration(ctx *Table_ttl_declarationContext) {
}

// ExitTable_ttl_declaration is called when production table_ttl_declaration is exited.
func (s *BaseClickHouseParserListener) ExitTable_ttl_declaration(ctx *Table_ttl_declarationContext) {}

// EnterNested_table is called when production nested_table is entered.
func (s *BaseClickHouseParserListener) EnterNested_table(ctx *Nested_tableContext) {}

// ExitNested_table is called when production nested_table is exited.
func (s *BaseClickHouseParserListener) ExitNested_table(ctx *Nested_tableContext) {}

// EnterName_type_pair_list is called when production name_type_pair_list is entered.
func (s *BaseClickHouseParserListener) EnterName_type_pair_list(ctx *Name_type_pair_listContext) {}

// ExitName_type_pair_list is called when production name_type_pair_list is exited.
func (s *BaseClickHouseParserListener) ExitName_type_pair_list(ctx *Name_type_pair_listContext) {}

// EnterName_type_pair is called when production name_type_pair is entered.
func (s *BaseClickHouseParserListener) EnterName_type_pair(ctx *Name_type_pairContext) {}

// ExitName_type_pair is called when production name_type_pair is exited.
func (s *BaseClickHouseParserListener) ExitName_type_pair(ctx *Name_type_pairContext) {}

// EnterCompound_name_type_pair is called when production compound_name_type_pair is entered.
func (s *BaseClickHouseParserListener) EnterCompound_name_type_pair(ctx *Compound_name_type_pairContext) {
}

// ExitCompound_name_type_pair is called when production compound_name_type_pair is exited.
func (s *BaseClickHouseParserListener) ExitCompound_name_type_pair(ctx *Compound_name_type_pairContext) {
}

// EnterColumn_declaration_list is called when production column_declaration_list is entered.
func (s *BaseClickHouseParserListener) EnterColumn_declaration_list(ctx *Column_declaration_listContext) {
}

// ExitColumn_declaration_list is called when production column_declaration_list is exited.
func (s *BaseClickHouseParserListener) ExitColumn_declaration_list(ctx *Column_declaration_listContext) {
}

// EnterColumn_declaration is called when production column_declaration is entered.
func (s *BaseClickHouseParserListener) EnterColumn_declaration(ctx *Column_declarationContext) {}

// ExitColumn_declaration is called when production column_declaration is exited.
func (s *BaseClickHouseParserListener) ExitColumn_declaration(ctx *Column_declarationContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseClickHouseParserListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseClickHouseParserListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterColumn_type is called when production column_type is entered.
func (s *BaseClickHouseParserListener) EnterColumn_type(ctx *Column_typeContext) {}

// ExitColumn_type is called when production column_type is exited.
func (s *BaseClickHouseParserListener) ExitColumn_type(ctx *Column_typeContext) {}

// EnterColumn_name_list is called when production column_name_list is entered.
func (s *BaseClickHouseParserListener) EnterColumn_name_list(ctx *Column_name_listContext) {}

// ExitColumn_name_list is called when production column_name_list is exited.
func (s *BaseClickHouseParserListener) ExitColumn_name_list(ctx *Column_name_listContext) {}

// EnterSelect_expr_list is called when production select_expr_list is entered.
func (s *BaseClickHouseParserListener) EnterSelect_expr_list(ctx *Select_expr_listContext) {}

// ExitSelect_expr_list is called when production select_expr_list is exited.
func (s *BaseClickHouseParserListener) ExitSelect_expr_list(ctx *Select_expr_listContext) {}

// EnterSelect_expr is called when production select_expr is entered.
func (s *BaseClickHouseParserListener) EnterSelect_expr(ctx *Select_exprContext) {}

// ExitSelect_expr is called when production select_expr is exited.
func (s *BaseClickHouseParserListener) ExitSelect_expr(ctx *Select_exprContext) {}

// EnterSelect_alias is called when production select_alias is entered.
func (s *BaseClickHouseParserListener) EnterSelect_alias(ctx *Select_aliasContext) {}

// ExitSelect_alias is called when production select_alias is exited.
func (s *BaseClickHouseParserListener) ExitSelect_alias(ctx *Select_aliasContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseClickHouseParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseClickHouseParserListener) ExitAlias(ctx *AliasContext) {}

// EnterAlias_name is called when production alias_name is entered.
func (s *BaseClickHouseParserListener) EnterAlias_name(ctx *Alias_nameContext) {}

// ExitAlias_name is called when production alias_name is exited.
func (s *BaseClickHouseParserListener) ExitAlias_name(ctx *Alias_nameContext) {}

// EnterTable_function is called when production table_function is entered.
func (s *BaseClickHouseParserListener) EnterTable_function(ctx *Table_functionContext) {}

// ExitTable_function is called when production table_function is exited.
func (s *BaseClickHouseParserListener) ExitTable_function(ctx *Table_functionContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseClickHouseParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseClickHouseParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterExpression_with_optional_alias is called when production expression_with_optional_alias is entered.
func (s *BaseClickHouseParserListener) EnterExpression_with_optional_alias(ctx *Expression_with_optional_aliasContext) {
}

// ExitExpression_with_optional_alias is called when production expression_with_optional_alias is exited.
func (s *BaseClickHouseParserListener) ExitExpression_with_optional_alias(ctx *Expression_with_optional_aliasContext) {
}

// EnterExprConcat is called when production ExprConcat is entered.
func (s *BaseClickHouseParserListener) EnterExprConcat(ctx *ExprConcatContext) {}

// ExitExprConcat is called when production ExprConcat is exited.
func (s *BaseClickHouseParserListener) ExitExprConcat(ctx *ExprConcatContext) {}

// EnterExprCase is called when production ExprCase is entered.
func (s *BaseClickHouseParserListener) EnterExprCase(ctx *ExprCaseContext) {}

// ExitExprCase is called when production ExprCase is exited.
func (s *BaseClickHouseParserListener) ExitExprCase(ctx *ExprCaseContext) {}

// EnterExprTupleElement is called when production ExprTupleElement is entered.
func (s *BaseClickHouseParserListener) EnterExprTupleElement(ctx *ExprTupleElementContext) {}

// ExitExprTupleElement is called when production ExprTupleElement is exited.
func (s *BaseClickHouseParserListener) ExitExprTupleElement(ctx *ExprTupleElementContext) {}

// EnterExprNot is called when production ExprNot is entered.
func (s *BaseClickHouseParserListener) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production ExprNot is exited.
func (s *BaseClickHouseParserListener) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprArray is called when production ExprArray is entered.
func (s *BaseClickHouseParserListener) EnterExprArray(ctx *ExprArrayContext) {}

// ExitExprArray is called when production ExprArray is exited.
func (s *BaseClickHouseParserListener) ExitExprArray(ctx *ExprArrayContext) {}

// EnterExprWithAlias is called when production ExprWithAlias is entered.
func (s *BaseClickHouseParserListener) EnterExprWithAlias(ctx *ExprWithAliasContext) {}

// ExitExprWithAlias is called when production ExprWithAlias is exited.
func (s *BaseClickHouseParserListener) ExitExprWithAlias(ctx *ExprWithAliasContext) {}

// EnterExprLogical is called when production ExprLogical is entered.
func (s *BaseClickHouseParserListener) EnterExprLogical(ctx *ExprLogicalContext) {}

// ExitExprLogical is called when production ExprLogical is exited.
func (s *BaseClickHouseParserListener) ExitExprLogical(ctx *ExprLogicalContext) {}

// EnterExprIn is called when production ExprIn is entered.
func (s *BaseClickHouseParserListener) EnterExprIn(ctx *ExprInContext) {}

// ExitExprIn is called when production ExprIn is exited.
func (s *BaseClickHouseParserListener) ExitExprIn(ctx *ExprInContext) {}

// EnterExprCast is called when production ExprCast is entered.
func (s *BaseClickHouseParserListener) EnterExprCast(ctx *ExprCastContext) {}

// ExitExprCast is called when production ExprCast is exited.
func (s *BaseClickHouseParserListener) ExitExprCast(ctx *ExprCastContext) {}

// EnterExprOr is called when production ExprOr is entered.
func (s *BaseClickHouseParserListener) EnterExprOr(ctx *ExprOrContext) {}

// ExitExprOr is called when production ExprOr is exited.
func (s *BaseClickHouseParserListener) ExitExprOr(ctx *ExprOrContext) {}

// EnterExprFunction is called when production ExprFunction is entered.
func (s *BaseClickHouseParserListener) EnterExprFunction(ctx *ExprFunctionContext) {}

// ExitExprFunction is called when production ExprFunction is exited.
func (s *BaseClickHouseParserListener) ExitExprFunction(ctx *ExprFunctionContext) {}

// EnterExprMul is called when production ExprMul is entered.
func (s *BaseClickHouseParserListener) EnterExprMul(ctx *ExprMulContext) {}

// ExitExprMul is called when production ExprMul is exited.
func (s *BaseClickHouseParserListener) ExitExprMul(ctx *ExprMulContext) {}

// EnterExprId is called when production ExprId is entered.
func (s *BaseClickHouseParserListener) EnterExprId(ctx *ExprIdContext) {}

// ExitExprId is called when production ExprId is exited.
func (s *BaseClickHouseParserListener) ExitExprId(ctx *ExprIdContext) {}

// EnterExprLambda is called when production ExprLambda is entered.
func (s *BaseClickHouseParserListener) EnterExprLambda(ctx *ExprLambdaContext) {}

// ExitExprLambda is called when production ExprLambda is exited.
func (s *BaseClickHouseParserListener) ExitExprLambda(ctx *ExprLambdaContext) {}

// EnterExprTernary is called when production ExprTernary is entered.
func (s *BaseClickHouseParserListener) EnterExprTernary(ctx *ExprTernaryContext) {}

// ExitExprTernary is called when production ExprTernary is exited.
func (s *BaseClickHouseParserListener) ExitExprTernary(ctx *ExprTernaryContext) {}

// EnterExprParen is called when production ExprParen is entered.
func (s *BaseClickHouseParserListener) EnterExprParen(ctx *ExprParenContext) {}

// ExitExprParen is called when production ExprParen is exited.
func (s *BaseClickHouseParserListener) ExitExprParen(ctx *ExprParenContext) {}

// EnterExprBetween is called when production ExprBetween is entered.
func (s *BaseClickHouseParserListener) EnterExprBetween(ctx *ExprBetweenContext) {}

// ExitExprBetween is called when production ExprBetween is exited.
func (s *BaseClickHouseParserListener) ExitExprBetween(ctx *ExprBetweenContext) {}

// EnterExprSubquery is called when production ExprSubquery is entered.
func (s *BaseClickHouseParserListener) EnterExprSubquery(ctx *ExprSubqueryContext) {}

// ExitExprSubquery is called when production ExprSubquery is exited.
func (s *BaseClickHouseParserListener) ExitExprSubquery(ctx *ExprSubqueryContext) {}

// EnterExprStar is called when production ExprStar is entered.
func (s *BaseClickHouseParserListener) EnterExprStar(ctx *ExprStarContext) {}

// ExitExprStar is called when production ExprStar is exited.
func (s *BaseClickHouseParserListener) ExitExprStar(ctx *ExprStarContext) {}

// EnterExprInterval is called when production ExprInterval is entered.
func (s *BaseClickHouseParserListener) EnterExprInterval(ctx *ExprIntervalContext) {}

// ExitExprInterval is called when production ExprInterval is exited.
func (s *BaseClickHouseParserListener) ExitExprInterval(ctx *ExprIntervalContext) {}

// EnterExprAnd is called when production ExprAnd is entered.
func (s *BaseClickHouseParserListener) EnterExprAnd(ctx *ExprAndContext) {}

// ExitExprAnd is called when production ExprAnd is exited.
func (s *BaseClickHouseParserListener) ExitExprAnd(ctx *ExprAndContext) {}

// EnterExprArrayElement is called when production ExprArrayElement is entered.
func (s *BaseClickHouseParserListener) EnterExprArrayElement(ctx *ExprArrayElementContext) {}

// ExitExprArrayElement is called when production ExprArrayElement is exited.
func (s *BaseClickHouseParserListener) ExitExprArrayElement(ctx *ExprArrayElementContext) {}

// EnterExprIsNull is called when production ExprIsNull is entered.
func (s *BaseClickHouseParserListener) EnterExprIsNull(ctx *ExprIsNullContext) {}

// ExitExprIsNull is called when production ExprIsNull is exited.
func (s *BaseClickHouseParserListener) ExitExprIsNull(ctx *ExprIsNullContext) {}

// EnterExprList is called when production ExprList is entered.
func (s *BaseClickHouseParserListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production ExprList is exited.
func (s *BaseClickHouseParserListener) ExitExprList(ctx *ExprListContext) {}

// EnterExprLiteral is called when production ExprLiteral is entered.
func (s *BaseClickHouseParserListener) EnterExprLiteral(ctx *ExprLiteralContext) {}

// ExitExprLiteral is called when production ExprLiteral is exited.
func (s *BaseClickHouseParserListener) ExitExprLiteral(ctx *ExprLiteralContext) {}

// EnterExprUnaryMinus is called when production ExprUnaryMinus is entered.
func (s *BaseClickHouseParserListener) EnterExprUnaryMinus(ctx *ExprUnaryMinusContext) {}

// ExitExprUnaryMinus is called when production ExprUnaryMinus is exited.
func (s *BaseClickHouseParserListener) ExitExprUnaryMinus(ctx *ExprUnaryMinusContext) {}

// EnterExprAdd is called when production ExprAdd is entered.
func (s *BaseClickHouseParserListener) EnterExprAdd(ctx *ExprAddContext) {}

// ExitExprAdd is called when production ExprAdd is exited.
func (s *BaseClickHouseParserListener) ExitExprAdd(ctx *ExprAddContext) {}

// EnterInterval_unit is called when production interval_unit is entered.
func (s *BaseClickHouseParserListener) EnterInterval_unit(ctx *Interval_unitContext) {}

// ExitInterval_unit is called when production interval_unit is exited.
func (s *BaseClickHouseParserListener) ExitInterval_unit(ctx *Interval_unitContext) {}

// EnterExpression_list is called when production expression_list is entered.
func (s *BaseClickHouseParserListener) EnterExpression_list(ctx *Expression_listContext) {}

// ExitExpression_list is called when production expression_list is exited.
func (s *BaseClickHouseParserListener) ExitExpression_list(ctx *Expression_listContext) {}

// EnterNot_empty_expression_list is called when production not_empty_expression_list is entered.
func (s *BaseClickHouseParserListener) EnterNot_empty_expression_list(ctx *Not_empty_expression_listContext) {
}

// ExitNot_empty_expression_list is called when production not_empty_expression_list is exited.
func (s *BaseClickHouseParserListener) ExitNot_empty_expression_list(ctx *Not_empty_expression_listContext) {
}

// EnterArray is called when production array is entered.
func (s *BaseClickHouseParserListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseClickHouseParserListener) ExitArray(ctx *ArrayContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseClickHouseParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseClickHouseParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunction_parameters is called when production function_parameters is entered.
func (s *BaseClickHouseParserListener) EnterFunction_parameters(ctx *Function_parametersContext) {}

// ExitFunction_parameters is called when production function_parameters is exited.
func (s *BaseClickHouseParserListener) ExitFunction_parameters(ctx *Function_parametersContext) {}

// EnterFunction_arguments is called when production function_arguments is entered.
func (s *BaseClickHouseParserListener) EnterFunction_arguments(ctx *Function_argumentsContext) {}

// ExitFunction_arguments is called when production function_arguments is exited.
func (s *BaseClickHouseParserListener) ExitFunction_arguments(ctx *Function_argumentsContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseClickHouseParserListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseClickHouseParserListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseClickHouseParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseClickHouseParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseClickHouseParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseClickHouseParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterCompound_identifier is called when production compound_identifier is entered.
func (s *BaseClickHouseParserListener) EnterCompound_identifier(ctx *Compound_identifierContext) {}

// ExitCompound_identifier is called when production compound_identifier is exited.
func (s *BaseClickHouseParserListener) ExitCompound_identifier(ctx *Compound_identifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseClickHouseParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseClickHouseParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterErr is called when production err is entered.
func (s *BaseClickHouseParserListener) EnterErr(ctx *ErrContext) {}

// ExitErr is called when production err is exited.
func (s *BaseClickHouseParserListener) ExitErr(ctx *ErrContext) {}
