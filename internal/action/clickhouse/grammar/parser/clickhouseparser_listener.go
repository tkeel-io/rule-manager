// Generated from ClickHouseParser.g4 by ANTLR 4.7.

package parser // ClickHouseParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ClickHouseParserListener is a complete listener for a parse tree produced by ClickHouseParser.
type ClickHouseParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSelect_query is called when entering the select_query production.
	EnterSelect_query(c *Select_queryContext)

	// EnterSelect_query_main is called when entering the select_query_main production.
	EnterSelect_query_main(c *Select_query_mainContext)

	// EnterSelect_with_step is called when entering the select_with_step production.
	EnterSelect_with_step(c *Select_with_stepContext)

	// EnterSelect_select_step is called when entering the select_select_step production.
	EnterSelect_select_step(c *Select_select_stepContext)

	// EnterSelect_from_step is called when entering the select_from_step production.
	EnterSelect_from_step(c *Select_from_stepContext)

	// EnterSelect_array_join_step is called when entering the select_array_join_step production.
	EnterSelect_array_join_step(c *Select_array_join_stepContext)

	// EnterSelect_sample_step is called when entering the select_sample_step production.
	EnterSelect_sample_step(c *Select_sample_stepContext)

	// EnterSample_ratio is called when entering the sample_ratio production.
	EnterSample_ratio(c *Sample_ratioContext)

	// EnterSelect_join_step is called when entering the select_join_step production.
	EnterSelect_join_step(c *Select_join_stepContext)

	// EnterSelect_join_right_part is called when entering the select_join_right_part production.
	EnterSelect_join_right_part(c *Select_join_right_partContext)

	// EnterSelect_prewhere_step is called when entering the select_prewhere_step production.
	EnterSelect_prewhere_step(c *Select_prewhere_stepContext)

	// EnterSelect_where_step is called when entering the select_where_step production.
	EnterSelect_where_step(c *Select_where_stepContext)

	// EnterSelect_groupby_step is called when entering the select_groupby_step production.
	EnterSelect_groupby_step(c *Select_groupby_stepContext)

	// EnterSelect_having_step is called when entering the select_having_step production.
	EnterSelect_having_step(c *Select_having_stepContext)

	// EnterSelect_orderby_step is called when entering the select_orderby_step production.
	EnterSelect_orderby_step(c *Select_orderby_stepContext)

	// EnterSelect_limit_step is called when entering the select_limit_step production.
	EnterSelect_limit_step(c *Select_limit_stepContext)

	// EnterSelect_limitby_step is called when entering the select_limitby_step production.
	EnterSelect_limitby_step(c *Select_limitby_stepContext)

	// EnterSettings_step is called when entering the settings_step production.
	EnterSettings_step(c *Settings_stepContext)

	// EnterSelect_format_step is called when entering the select_format_step production.
	EnterSelect_format_step(c *Select_format_stepContext)

	// EnterInsert_query is called when entering the insert_query production.
	EnterInsert_query(c *Insert_queryContext)

	// EnterCreate_query is called when entering the create_query production.
	EnterCreate_query(c *Create_queryContext)

	// EnterRename_query is called when entering the rename_query production.
	EnterRename_query(c *Rename_queryContext)

	// EnterDrop_query is called when entering the drop_query production.
	EnterDrop_query(c *Drop_queryContext)

	// EnterAlter_query is called when entering the alter_query production.
	EnterAlter_query(c *Alter_queryContext)

	// EnterAlter_query_element is called when entering the alter_query_element production.
	EnterAlter_query_element(c *Alter_query_elementContext)

	// EnterClickhouse_type is called when entering the clickhouse_type production.
	EnterClickhouse_type(c *Clickhouse_typeContext)

	// EnterSimple_type is called when entering the simple_type production.
	EnterSimple_type(c *Simple_typeContext)

	// EnterEnum_entry is called when entering the enum_entry production.
	EnterEnum_entry(c *Enum_entryContext)

	// EnterUse_query is called when entering the use_query production.
	EnterUse_query(c *Use_queryContext)

	// EnterSet_query is called when entering the set_query production.
	EnterSet_query(c *Set_queryContext)

	// EnterAssignment_list is called when entering the assignment_list production.
	EnterAssignment_list(c *Assignment_listContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterKill_query_query is called when entering the kill_query_query production.
	EnterKill_query_query(c *Kill_query_queryContext)

	// EnterOptimize_query is called when entering the optimize_query production.
	EnterOptimize_query(c *Optimize_queryContext)

	// EnterTable_properties_query is called when entering the table_properties_query production.
	EnterTable_properties_query(c *Table_properties_queryContext)

	// EnterShow_tables_query is called when entering the show_tables_query production.
	EnterShow_tables_query(c *Show_tables_queryContext)

	// EnterShow_processlist_query is called when entering the show_processlist_query production.
	EnterShow_processlist_query(c *Show_processlist_queryContext)

	// EnterCheck_query is called when entering the check_query production.
	EnterCheck_query(c *Check_queryContext)

	// EnterFull_table_name is called when entering the full_table_name production.
	EnterFull_table_name(c *Full_table_nameContext)

	// EnterPartition_name is called when entering the partition_name production.
	EnterPartition_name(c *Partition_nameContext)

	// EnterCluster_name is called when entering the cluster_name production.
	EnterCluster_name(c *Cluster_nameContext)

	// EnterDatabase_name is called when entering the database_name production.
	EnterDatabase_name(c *Database_nameContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterFormat_name is called when entering the format_name production.
	EnterFormat_name(c *Format_nameContext)

	// EnterQuery_outfile_step is called when entering the query_outfile_step production.
	EnterQuery_outfile_step(c *Query_outfile_stepContext)

	// EnterEngine is called when entering the engine production.
	EnterEngine(c *EngineContext)

	// EnterIdentifier_with_optional_parameters is called when entering the identifier_with_optional_parameters production.
	EnterIdentifier_with_optional_parameters(c *Identifier_with_optional_parametersContext)

	// EnterIdentifier_with_parameters is called when entering the identifier_with_parameters production.
	EnterIdentifier_with_parameters(c *Identifier_with_parametersContext)

	// EnterOrder_by_expression_list is called when entering the order_by_expression_list production.
	EnterOrder_by_expression_list(c *Order_by_expression_listContext)

	// EnterOrder_by_element is called when entering the order_by_element production.
	EnterOrder_by_element(c *Order_by_elementContext)

	// EnterTable_ttl_list is called when entering the table_ttl_list production.
	EnterTable_ttl_list(c *Table_ttl_listContext)

	// EnterPartition_by_element is called when entering the partition_by_element production.
	EnterPartition_by_element(c *Partition_by_elementContext)

	// EnterTable_ttl_declaration is called when entering the table_ttl_declaration production.
	EnterTable_ttl_declaration(c *Table_ttl_declarationContext)

	// EnterNested_table is called when entering the nested_table production.
	EnterNested_table(c *Nested_tableContext)

	// EnterName_type_pair_list is called when entering the name_type_pair_list production.
	EnterName_type_pair_list(c *Name_type_pair_listContext)

	// EnterName_type_pair is called when entering the name_type_pair production.
	EnterName_type_pair(c *Name_type_pairContext)

	// EnterCompound_name_type_pair is called when entering the compound_name_type_pair production.
	EnterCompound_name_type_pair(c *Compound_name_type_pairContext)

	// EnterColumn_declaration_list is called when entering the column_declaration_list production.
	EnterColumn_declaration_list(c *Column_declaration_listContext)

	// EnterColumn_declaration is called when entering the column_declaration production.
	EnterColumn_declaration(c *Column_declarationContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterColumn_type is called when entering the column_type production.
	EnterColumn_type(c *Column_typeContext)

	// EnterColumn_name_list is called when entering the column_name_list production.
	EnterColumn_name_list(c *Column_name_listContext)

	// EnterSelect_expr_list is called when entering the select_expr_list production.
	EnterSelect_expr_list(c *Select_expr_listContext)

	// EnterSelect_expr is called when entering the select_expr production.
	EnterSelect_expr(c *Select_exprContext)

	// EnterSelect_alias is called when entering the select_alias production.
	EnterSelect_alias(c *Select_aliasContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterAlias_name is called when entering the alias_name production.
	EnterAlias_name(c *Alias_nameContext)

	// EnterTable_function is called when entering the table_function production.
	EnterTable_function(c *Table_functionContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterExpression_with_optional_alias is called when entering the expression_with_optional_alias production.
	EnterExpression_with_optional_alias(c *Expression_with_optional_aliasContext)

	// EnterExprConcat is called when entering the ExprConcat production.
	EnterExprConcat(c *ExprConcatContext)

	// EnterExprCase is called when entering the ExprCase production.
	EnterExprCase(c *ExprCaseContext)

	// EnterExprTupleElement is called when entering the ExprTupleElement production.
	EnterExprTupleElement(c *ExprTupleElementContext)

	// EnterExprNot is called when entering the ExprNot production.
	EnterExprNot(c *ExprNotContext)

	// EnterExprArray is called when entering the ExprArray production.
	EnterExprArray(c *ExprArrayContext)

	// EnterExprWithAlias is called when entering the ExprWithAlias production.
	EnterExprWithAlias(c *ExprWithAliasContext)

	// EnterExprLogical is called when entering the ExprLogical production.
	EnterExprLogical(c *ExprLogicalContext)

	// EnterExprIn is called when entering the ExprIn production.
	EnterExprIn(c *ExprInContext)

	// EnterExprCast is called when entering the ExprCast production.
	EnterExprCast(c *ExprCastContext)

	// EnterExprOr is called when entering the ExprOr production.
	EnterExprOr(c *ExprOrContext)

	// EnterExprFunction is called when entering the ExprFunction production.
	EnterExprFunction(c *ExprFunctionContext)

	// EnterExprMul is called when entering the ExprMul production.
	EnterExprMul(c *ExprMulContext)

	// EnterExprId is called when entering the ExprId production.
	EnterExprId(c *ExprIdContext)

	// EnterExprLambda is called when entering the ExprLambda production.
	EnterExprLambda(c *ExprLambdaContext)

	// EnterExprTernary is called when entering the ExprTernary production.
	EnterExprTernary(c *ExprTernaryContext)

	// EnterExprParen is called when entering the ExprParen production.
	EnterExprParen(c *ExprParenContext)

	// EnterExprBetween is called when entering the ExprBetween production.
	EnterExprBetween(c *ExprBetweenContext)

	// EnterExprSubquery is called when entering the ExprSubquery production.
	EnterExprSubquery(c *ExprSubqueryContext)

	// EnterExprStar is called when entering the ExprStar production.
	EnterExprStar(c *ExprStarContext)

	// EnterExprInterval is called when entering the ExprInterval production.
	EnterExprInterval(c *ExprIntervalContext)

	// EnterExprAnd is called when entering the ExprAnd production.
	EnterExprAnd(c *ExprAndContext)

	// EnterExprArrayElement is called when entering the ExprArrayElement production.
	EnterExprArrayElement(c *ExprArrayElementContext)

	// EnterExprIsNull is called when entering the ExprIsNull production.
	EnterExprIsNull(c *ExprIsNullContext)

	// EnterExprList is called when entering the ExprList production.
	EnterExprList(c *ExprListContext)

	// EnterExprLiteral is called when entering the ExprLiteral production.
	EnterExprLiteral(c *ExprLiteralContext)

	// EnterExprUnaryMinus is called when entering the ExprUnaryMinus production.
	EnterExprUnaryMinus(c *ExprUnaryMinusContext)

	// EnterExprAdd is called when entering the ExprAdd production.
	EnterExprAdd(c *ExprAddContext)

	// EnterInterval_unit is called when entering the interval_unit production.
	EnterInterval_unit(c *Interval_unitContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterNot_empty_expression_list is called when entering the not_empty_expression_list production.
	EnterNot_empty_expression_list(c *Not_empty_expression_listContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFunction_parameters is called when entering the function_parameters production.
	EnterFunction_parameters(c *Function_parametersContext)

	// EnterFunction_arguments is called when entering the function_arguments production.
	EnterFunction_arguments(c *Function_argumentsContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterCompound_identifier is called when entering the compound_identifier production.
	EnterCompound_identifier(c *Compound_identifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterErr is called when entering the err production.
	EnterErr(c *ErrContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSelect_query is called when exiting the select_query production.
	ExitSelect_query(c *Select_queryContext)

	// ExitSelect_query_main is called when exiting the select_query_main production.
	ExitSelect_query_main(c *Select_query_mainContext)

	// ExitSelect_with_step is called when exiting the select_with_step production.
	ExitSelect_with_step(c *Select_with_stepContext)

	// ExitSelect_select_step is called when exiting the select_select_step production.
	ExitSelect_select_step(c *Select_select_stepContext)

	// ExitSelect_from_step is called when exiting the select_from_step production.
	ExitSelect_from_step(c *Select_from_stepContext)

	// ExitSelect_array_join_step is called when exiting the select_array_join_step production.
	ExitSelect_array_join_step(c *Select_array_join_stepContext)

	// ExitSelect_sample_step is called when exiting the select_sample_step production.
	ExitSelect_sample_step(c *Select_sample_stepContext)

	// ExitSample_ratio is called when exiting the sample_ratio production.
	ExitSample_ratio(c *Sample_ratioContext)

	// ExitSelect_join_step is called when exiting the select_join_step production.
	ExitSelect_join_step(c *Select_join_stepContext)

	// ExitSelect_join_right_part is called when exiting the select_join_right_part production.
	ExitSelect_join_right_part(c *Select_join_right_partContext)

	// ExitSelect_prewhere_step is called when exiting the select_prewhere_step production.
	ExitSelect_prewhere_step(c *Select_prewhere_stepContext)

	// ExitSelect_where_step is called when exiting the select_where_step production.
	ExitSelect_where_step(c *Select_where_stepContext)

	// ExitSelect_groupby_step is called when exiting the select_groupby_step production.
	ExitSelect_groupby_step(c *Select_groupby_stepContext)

	// ExitSelect_having_step is called when exiting the select_having_step production.
	ExitSelect_having_step(c *Select_having_stepContext)

	// ExitSelect_orderby_step is called when exiting the select_orderby_step production.
	ExitSelect_orderby_step(c *Select_orderby_stepContext)

	// ExitSelect_limit_step is called when exiting the select_limit_step production.
	ExitSelect_limit_step(c *Select_limit_stepContext)

	// ExitSelect_limitby_step is called when exiting the select_limitby_step production.
	ExitSelect_limitby_step(c *Select_limitby_stepContext)

	// ExitSettings_step is called when exiting the settings_step production.
	ExitSettings_step(c *Settings_stepContext)

	// ExitSelect_format_step is called when exiting the select_format_step production.
	ExitSelect_format_step(c *Select_format_stepContext)

	// ExitInsert_query is called when exiting the insert_query production.
	ExitInsert_query(c *Insert_queryContext)

	// ExitCreate_query is called when exiting the create_query production.
	ExitCreate_query(c *Create_queryContext)

	// ExitRename_query is called when exiting the rename_query production.
	ExitRename_query(c *Rename_queryContext)

	// ExitDrop_query is called when exiting the drop_query production.
	ExitDrop_query(c *Drop_queryContext)

	// ExitAlter_query is called when exiting the alter_query production.
	ExitAlter_query(c *Alter_queryContext)

	// ExitAlter_query_element is called when exiting the alter_query_element production.
	ExitAlter_query_element(c *Alter_query_elementContext)

	// ExitClickhouse_type is called when exiting the clickhouse_type production.
	ExitClickhouse_type(c *Clickhouse_typeContext)

	// ExitSimple_type is called when exiting the simple_type production.
	ExitSimple_type(c *Simple_typeContext)

	// ExitEnum_entry is called when exiting the enum_entry production.
	ExitEnum_entry(c *Enum_entryContext)

	// ExitUse_query is called when exiting the use_query production.
	ExitUse_query(c *Use_queryContext)

	// ExitSet_query is called when exiting the set_query production.
	ExitSet_query(c *Set_queryContext)

	// ExitAssignment_list is called when exiting the assignment_list production.
	ExitAssignment_list(c *Assignment_listContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitKill_query_query is called when exiting the kill_query_query production.
	ExitKill_query_query(c *Kill_query_queryContext)

	// ExitOptimize_query is called when exiting the optimize_query production.
	ExitOptimize_query(c *Optimize_queryContext)

	// ExitTable_properties_query is called when exiting the table_properties_query production.
	ExitTable_properties_query(c *Table_properties_queryContext)

	// ExitShow_tables_query is called when exiting the show_tables_query production.
	ExitShow_tables_query(c *Show_tables_queryContext)

	// ExitShow_processlist_query is called when exiting the show_processlist_query production.
	ExitShow_processlist_query(c *Show_processlist_queryContext)

	// ExitCheck_query is called when exiting the check_query production.
	ExitCheck_query(c *Check_queryContext)

	// ExitFull_table_name is called when exiting the full_table_name production.
	ExitFull_table_name(c *Full_table_nameContext)

	// ExitPartition_name is called when exiting the partition_name production.
	ExitPartition_name(c *Partition_nameContext)

	// ExitCluster_name is called when exiting the cluster_name production.
	ExitCluster_name(c *Cluster_nameContext)

	// ExitDatabase_name is called when exiting the database_name production.
	ExitDatabase_name(c *Database_nameContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitFormat_name is called when exiting the format_name production.
	ExitFormat_name(c *Format_nameContext)

	// ExitQuery_outfile_step is called when exiting the query_outfile_step production.
	ExitQuery_outfile_step(c *Query_outfile_stepContext)

	// ExitEngine is called when exiting the engine production.
	ExitEngine(c *EngineContext)

	// ExitIdentifier_with_optional_parameters is called when exiting the identifier_with_optional_parameters production.
	ExitIdentifier_with_optional_parameters(c *Identifier_with_optional_parametersContext)

	// ExitIdentifier_with_parameters is called when exiting the identifier_with_parameters production.
	ExitIdentifier_with_parameters(c *Identifier_with_parametersContext)

	// ExitOrder_by_expression_list is called when exiting the order_by_expression_list production.
	ExitOrder_by_expression_list(c *Order_by_expression_listContext)

	// ExitOrder_by_element is called when exiting the order_by_element production.
	ExitOrder_by_element(c *Order_by_elementContext)

	// ExitTable_ttl_list is called when exiting the table_ttl_list production.
	ExitTable_ttl_list(c *Table_ttl_listContext)

	// ExitPartition_by_element is called when exiting the partition_by_element production.
	ExitPartition_by_element(c *Partition_by_elementContext)

	// ExitTable_ttl_declaration is called when exiting the table_ttl_declaration production.
	ExitTable_ttl_declaration(c *Table_ttl_declarationContext)

	// ExitNested_table is called when exiting the nested_table production.
	ExitNested_table(c *Nested_tableContext)

	// ExitName_type_pair_list is called when exiting the name_type_pair_list production.
	ExitName_type_pair_list(c *Name_type_pair_listContext)

	// ExitName_type_pair is called when exiting the name_type_pair production.
	ExitName_type_pair(c *Name_type_pairContext)

	// ExitCompound_name_type_pair is called when exiting the compound_name_type_pair production.
	ExitCompound_name_type_pair(c *Compound_name_type_pairContext)

	// ExitColumn_declaration_list is called when exiting the column_declaration_list production.
	ExitColumn_declaration_list(c *Column_declaration_listContext)

	// ExitColumn_declaration is called when exiting the column_declaration production.
	ExitColumn_declaration(c *Column_declarationContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitColumn_type is called when exiting the column_type production.
	ExitColumn_type(c *Column_typeContext)

	// ExitColumn_name_list is called when exiting the column_name_list production.
	ExitColumn_name_list(c *Column_name_listContext)

	// ExitSelect_expr_list is called when exiting the select_expr_list production.
	ExitSelect_expr_list(c *Select_expr_listContext)

	// ExitSelect_expr is called when exiting the select_expr production.
	ExitSelect_expr(c *Select_exprContext)

	// ExitSelect_alias is called when exiting the select_alias production.
	ExitSelect_alias(c *Select_aliasContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitAlias_name is called when exiting the alias_name production.
	ExitAlias_name(c *Alias_nameContext)

	// ExitTable_function is called when exiting the table_function production.
	ExitTable_function(c *Table_functionContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitExpression_with_optional_alias is called when exiting the expression_with_optional_alias production.
	ExitExpression_with_optional_alias(c *Expression_with_optional_aliasContext)

	// ExitExprConcat is called when exiting the ExprConcat production.
	ExitExprConcat(c *ExprConcatContext)

	// ExitExprCase is called when exiting the ExprCase production.
	ExitExprCase(c *ExprCaseContext)

	// ExitExprTupleElement is called when exiting the ExprTupleElement production.
	ExitExprTupleElement(c *ExprTupleElementContext)

	// ExitExprNot is called when exiting the ExprNot production.
	ExitExprNot(c *ExprNotContext)

	// ExitExprArray is called when exiting the ExprArray production.
	ExitExprArray(c *ExprArrayContext)

	// ExitExprWithAlias is called when exiting the ExprWithAlias production.
	ExitExprWithAlias(c *ExprWithAliasContext)

	// ExitExprLogical is called when exiting the ExprLogical production.
	ExitExprLogical(c *ExprLogicalContext)

	// ExitExprIn is called when exiting the ExprIn production.
	ExitExprIn(c *ExprInContext)

	// ExitExprCast is called when exiting the ExprCast production.
	ExitExprCast(c *ExprCastContext)

	// ExitExprOr is called when exiting the ExprOr production.
	ExitExprOr(c *ExprOrContext)

	// ExitExprFunction is called when exiting the ExprFunction production.
	ExitExprFunction(c *ExprFunctionContext)

	// ExitExprMul is called when exiting the ExprMul production.
	ExitExprMul(c *ExprMulContext)

	// ExitExprId is called when exiting the ExprId production.
	ExitExprId(c *ExprIdContext)

	// ExitExprLambda is called when exiting the ExprLambda production.
	ExitExprLambda(c *ExprLambdaContext)

	// ExitExprTernary is called when exiting the ExprTernary production.
	ExitExprTernary(c *ExprTernaryContext)

	// ExitExprParen is called when exiting the ExprParen production.
	ExitExprParen(c *ExprParenContext)

	// ExitExprBetween is called when exiting the ExprBetween production.
	ExitExprBetween(c *ExprBetweenContext)

	// ExitExprSubquery is called when exiting the ExprSubquery production.
	ExitExprSubquery(c *ExprSubqueryContext)

	// ExitExprStar is called when exiting the ExprStar production.
	ExitExprStar(c *ExprStarContext)

	// ExitExprInterval is called when exiting the ExprInterval production.
	ExitExprInterval(c *ExprIntervalContext)

	// ExitExprAnd is called when exiting the ExprAnd production.
	ExitExprAnd(c *ExprAndContext)

	// ExitExprArrayElement is called when exiting the ExprArrayElement production.
	ExitExprArrayElement(c *ExprArrayElementContext)

	// ExitExprIsNull is called when exiting the ExprIsNull production.
	ExitExprIsNull(c *ExprIsNullContext)

	// ExitExprList is called when exiting the ExprList production.
	ExitExprList(c *ExprListContext)

	// ExitExprLiteral is called when exiting the ExprLiteral production.
	ExitExprLiteral(c *ExprLiteralContext)

	// ExitExprUnaryMinus is called when exiting the ExprUnaryMinus production.
	ExitExprUnaryMinus(c *ExprUnaryMinusContext)

	// ExitExprAdd is called when exiting the ExprAdd production.
	ExitExprAdd(c *ExprAddContext)

	// ExitInterval_unit is called when exiting the interval_unit production.
	ExitInterval_unit(c *Interval_unitContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitNot_empty_expression_list is called when exiting the not_empty_expression_list production.
	ExitNot_empty_expression_list(c *Not_empty_expression_listContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFunction_parameters is called when exiting the function_parameters production.
	ExitFunction_parameters(c *Function_parametersContext)

	// ExitFunction_arguments is called when exiting the function_arguments production.
	ExitFunction_arguments(c *Function_argumentsContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitCompound_identifier is called when exiting the compound_identifier production.
	ExitCompound_identifier(c *Compound_identifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitErr is called when exiting the err production.
	ExitErr(c *ErrContext)
}
