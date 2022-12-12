// Generated from java-escape by ANTLR 4.11.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ConfigParser}.
 */
public interface ConfigListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ConfigParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterStat(ConfigParser.StatContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConfigParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitStat(ConfigParser.StatContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConfigParser#sections}.
	 * @param ctx the parse tree
	 */
	void enterSections(ConfigParser.SectionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConfigParser#sections}.
	 * @param ctx the parse tree
	 */
	void exitSections(ConfigParser.SectionsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code GetSection}
	 * labeled alternative in {@link ConfigParser#section}.
	 * @param ctx the parse tree
	 */
	void enterGetSection(ConfigParser.GetSectionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code GetSection}
	 * labeled alternative in {@link ConfigParser#section}.
	 * @param ctx the parse tree
	 */
	void exitGetSection(ConfigParser.GetSectionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConfigParser#fields}.
	 * @param ctx the parse tree
	 */
	void enterFields(ConfigParser.FieldsContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConfigParser#fields}.
	 * @param ctx the parse tree
	 */
	void exitFields(ConfigParser.FieldsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code GetField}
	 * labeled alternative in {@link ConfigParser#field}.
	 * @param ctx the parse tree
	 */
	void enterGetField(ConfigParser.GetFieldContext ctx);
	/**
	 * Exit a parse tree produced by the {@code GetField}
	 * labeled alternative in {@link ConfigParser#field}.
	 * @param ctx the parse tree
	 */
	void exitGetField(ConfigParser.GetFieldContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConfigParser#fieldName}.
	 * @param ctx the parse tree
	 */
	void enterFieldName(ConfigParser.FieldNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConfigParser#fieldName}.
	 * @param ctx the parse tree
	 */
	void exitFieldName(ConfigParser.FieldNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link ConfigParser#fieldValue}.
	 * @param ctx the parse tree
	 */
	void enterFieldValue(ConfigParser.FieldValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link ConfigParser#fieldValue}.
	 * @param ctx the parse tree
	 */
	void exitFieldValue(ConfigParser.FieldValueContext ctx);
}