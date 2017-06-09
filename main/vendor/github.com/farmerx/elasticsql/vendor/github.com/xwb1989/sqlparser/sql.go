//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr

	/*
	   for CreateTable
	*/
	createTableStmt   CreateTable
	columnDefinition  *ColumnDefinition
	columnDefinitions ColumnDefinitions
	columnAtts        ColumnAtts
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const LIST_ARG = 57381
const COMMENT = 57382
const LE = 57383
const GE = 57384
const NE = 57385
const NULL_SAFE_EQUAL = 57386
const PRIMARY = 57387
const UNIQUE = 57388
const UNION = 57389
const MINUS = 57390
const EXCEPT = 57391
const INTERSECT = 57392
const JOIN = 57393
const STRAIGHT_JOIN = 57394
const LEFT = 57395
const RIGHT = 57396
const INNER = 57397
const OUTER = 57398
const CROSS = 57399
const NATURAL = 57400
const USE = 57401
const FORCE = 57402
const ON = 57403
const OR = 57404
const AND = 57405
const NOT = 57406
const UNARY = 57407
const CASE = 57408
const WHEN = 57409
const THEN = 57410
const ELSE = 57411
const END = 57412
const CREATE = 57413
const ALTER = 57414
const DROP = 57415
const RENAME = 57416
const ANALYZE = 57417
const TABLE = 57418
const INDEX = 57419
const VIEW = 57420
const TO = 57421
const IGNORE = 57422
const IF = 57423
const USING = 57424
const SHOW = 57425
const DESCRIBE = 57426
const EXPLAIN = 57427
const BIT = 57428
const TINYINT = 57429
const SMALLINT = 57430
const MEDIUMINT = 57431
const INT = 57432
const INTEGER = 57433
const BIGINT = 57434
const REAL = 57435
const DOUBLE = 57436
const FLOAT = 57437
const UNSIGNED = 57438
const ZEROFILL = 57439
const DECIMAL = 57440
const NUMERIC = 57441
const DATE = 57442
const TIME = 57443
const TIMESTAMP = 57444
const DATETIME = 57445
const YEAR = 57446
const TEXT = 57447
const CHAR = 57448
const VARCHAR = 57449
const NULLX = 57450
const AUTO_INCREMENT = 57451
const BOOL = 57452
const APPROXNUM = 57453
const INTNUM = 57454

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"'('",
	"'='",
	"'<'",
	"'>'",
	"'~'",
	"PRIMARY",
	"UNIQUE",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"','",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	"'&'",
	"'|'",
	"'^'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'.'",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"BIT",
	"TINYINT",
	"SMALLINT",
	"MEDIUMINT",
	"INT",
	"INTEGER",
	"BIGINT",
	"REAL",
	"DOUBLE",
	"FLOAT",
	"UNSIGNED",
	"ZEROFILL",
	"DECIMAL",
	"NUMERIC",
	"DATE",
	"TIME",
	"TIMESTAMP",
	"DATETIME",
	"YEAR",
	"TEXT",
	"CHAR",
	"VARCHAR",
	"NULLX",
	"AUTO_INCREMENT",
	"BOOL",
	"APPROXNUM",
	"INTNUM",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 254
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 689

var yyAct = [...]int{

	95, 294, 159, 429, 92, 358, 246, 93, 63, 162,
	368, 250, 364, 197, 286, 103, 237, 208, 439, 86,
	81, 177, 161, 3, 260, 261, 262, 263, 264, 82,
	265, 266, 439, 51, 439, 29, 30, 31, 32, 66,
	136, 135, 71, 65, 442, 74, 426, 64, 407, 78,
	427, 54, 130, 229, 297, 292, 130, 130, 397, 87,
	52, 53, 363, 185, 404, 398, 229, 77, 339, 341,
	69, 120, 44, 386, 45, 47, 48, 49, 385, 256,
	128, 124, 384, 403, 405, 132, 271, 70, 238, 349,
	441, 134, 117, 163, 113, 343, 73, 164, 340, 42,
	227, 72, 50, 396, 440, 121, 438, 46, 123, 158,
	160, 228, 171, 66, 168, 119, 66, 65, 181, 180,
	65, 175, 426, 381, 348, 345, 298, 291, 281, 279,
	148, 149, 150, 87, 203, 181, 217, 135, 230, 39,
	207, 41, 204, 215, 216, 201, 219, 220, 221, 222,
	223, 224, 225, 226, 210, 205, 206, 399, 136, 135,
	202, 115, 195, 179, 238, 287, 284, 252, 231, 87,
	87, 191, 287, 351, 66, 66, 136, 135, 65, 244,
	127, 218, 242, 383, 333, 382, 253, 233, 235, 334,
	189, 60, 337, 192, 336, 241, 335, 245, 248, 146,
	147, 148, 149, 150, 234, 331, 98, 115, 229, 427,
	332, 102, 231, 270, 108, 201, 274, 275, 257, 272,
	254, 85, 99, 100, 101, 178, 392, 353, 210, 116,
	273, 90, 278, 200, 178, 106, 129, 87, 76, 29,
	30, 31, 32, 199, 211, 188, 190, 187, 289, 417,
	209, 111, 293, 283, 114, 280, 89, 290, 285, 416,
	104, 105, 83, 14, 15, 16, 17, 109, 415, 173,
	165, 258, 329, 330, 201, 201, 394, 395, 374, 347,
	115, 174, 107, 130, 369, 391, 365, 350, 79, 182,
	14, 18, 169, 66, 167, 355, 166, 354, 356, 359,
	143, 144, 145, 146, 147, 148, 149, 150, 110, 360,
	422, 423, 434, 410, 232, 409, 408, 72, 366, 367,
	200, 143, 144, 145, 146, 147, 148, 149, 150, 67,
	199, 251, 370, 371, 372, 375, 373, 376, 344, 269,
	342, 326, 325, 194, 20, 21, 23, 22, 24, 387,
	193, 176, 125, 122, 388, 268, 25, 26, 27, 377,
	346, 390, 143, 144, 145, 146, 147, 148, 149, 150,
	276, 133, 143, 144, 145, 146, 147, 148, 149, 150,
	118, 61, 436, 231, 80, 411, 75, 72, 112, 425,
	413, 424, 389, 352, 419, 359, 59, 421, 420, 412,
	437, 414, 143, 144, 145, 146, 147, 148, 149, 150,
	14, 277, 212, 428, 213, 214, 430, 430, 430, 66,
	431, 432, 444, 65, 183, 33, 126, 433, 260, 261,
	262, 263, 264, 240, 265, 266, 57, 55, 295, 380,
	445, 35, 36, 37, 38, 446, 296, 447, 313, 314,
	315, 316, 317, 318, 319, 320, 321, 322, 247, 379,
	323, 324, 308, 309, 310, 311, 312, 307, 305, 306,
	98, 328, 178, 62, 14, 102, 14, 443, 108, 418,
	34, 402, 401, 361, 302, 85, 99, 100, 101, 98,
	304, 303, 400, 406, 102, 90, 362, 108, 300, 106,
	301, 19, 249, 299, 67, 99, 100, 101, 184, 40,
	255, 186, 43, 68, 90, 243, 172, 435, 106, 393,
	89, 357, 378, 327, 104, 105, 83, 282, 170, 236,
	97, 109, 94, 96, 288, 91, 239, 98, 14, 89,
	137, 88, 102, 104, 105, 108, 107, 338, 198, 259,
	109, 196, 67, 99, 100, 101, 84, 267, 102, 131,
	56, 108, 90, 28, 58, 107, 106, 13, 67, 99,
	100, 101, 12, 11, 10, 9, 8, 7, 165, 6,
	5, 4, 106, 2, 1, 0, 0, 89, 0, 0,
	0, 104, 105, 0, 0, 0, 0, 0, 109, 102,
	0, 0, 108, 0, 0, 0, 0, 104, 105, 67,
	99, 100, 101, 107, 109, 0, 0, 0, 0, 165,
	0, 0, 0, 106, 0, 0, 0, 0, 0, 107,
	0, 138, 142, 140, 141, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 104, 105,
	0, 154, 155, 156, 157, 109, 151, 152, 153, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	107, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	139, 143, 144, 145, 146, 147, 148, 149, 150,
}
var yyPact = [...]int{

	258, -1000, -1000, 187, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	48, -21, 16, -16, 11, -1000, -1000, -1000, 471, 420,
	-1000, -1000, -1000, 418, -1000, 367, 346, 464, 294, -26,
	-5, 282, -1000, 5, 282, -1000, 351, -29, 282, -29,
	349, -1000, -1000, -1000, -1000, -1000, 450, -1000, 268, 346,
	355, 15, 346, 151, -1000, 183, -1000, 13, 345, 45,
	282, -1000, -1000, 318, -1000, -13, 317, 406, 113, 282,
	-1000, 227, -1000, -1000, 352, 12, 108, 610, -1000, 517,
	469, -1000, -1000, -1000, 574, 251, 249, -1000, 247, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 574,
	-1000, 236, 294, 316, 462, 294, 574, 282, 244, 404,
	-34, -1000, 158, -1000, 315, -1000, -1000, 308, -1000, 198,
	450, -1000, -1000, 282, 66, 517, 517, 574, 205, 391,
	574, 574, 111, 574, 574, 574, 574, 574, 574, 574,
	574, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 610,
	-28, -17, 10, 610, -1000, 533, 186, 450, -1000, 471,
	6, 331, 405, 294, 294, 224, -1000, 445, 517, -1000,
	331, -1000, 296, -1000, 100, 282, -1000, -15, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 215, 371, 320, 285,
	7, -1000, -1000, -1000, -1000, -1000, 68, 331, -1000, 533,
	-1000, -1000, 205, 574, 574, 331, 301, -1000, 386, 125,
	125, 125, 54, 54, -1000, -1000, -1000, -1000, -1000, 574,
	-1000, 331, -1000, 1, 450, 0, 82, -1000, 517, 98,
	225, 187, 105, -1, -1000, 445, 423, 432, 108, -2,
	-1000, 347, 307, -1000, -1000, 306, -1000, 460, 198, 198,
	-1000, -1000, 148, 127, 139, 137, 135, 3, -1000, 305,
	-33, 303, -3, -1000, 331, 291, 574, -1000, 331, -1000,
	-4, -1000, 4, -1000, 574, 90, -1000, 363, 171, -1000,
	-1000, -1000, 294, 423, -1000, 574, 574, 296, -1000, -1000,
	-49, -1000, -1000, 241, -1000, 241, 241, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	239, 239, 239, 233, 233, -1000, -1000, 447, 425, 371,
	56, -1000, 128, -1000, 126, -1000, -1000, -1000, -1000, -10,
	-14, -19, -1000, -1000, -1000, -1000, 574, 331, -1000, -1000,
	331, 574, 361, 225, -1000, -1000, 229, 170, -1000, 250,
	-1000, 33, -64, -1000, -1000, 279, -1000, -1000, -1000, 278,
	-1000, -1000, -1000, -1000, 276, -1000, -1000, -1000, 445, 517,
	574, 517, -1000, -1000, 223, 214, 204, 331, 331, 472,
	-1000, 574, 574, -1000, -1000, -1000, 372, -1000, 274, -1000,
	-1000, -1000, -1000, 360, -1000, 358, -1000, -1000, -82, 153,
	-6, 423, 108, 152, 108, 282, 282, 282, 294, 331,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 275, 366, -22,
	-1000, -24, -38, 151, -84, -1000, 470, 401, -1000, 282,
	-1000, -1000, -1000, -1000, 282, -1000, 282, -1000,
}
var yyPgo = [...]int{

	0, 584, 583, 22, 581, 580, 579, 577, 576, 575,
	574, 573, 572, 567, 425, 564, 563, 560, 20, 29,
	559, 557, 556, 551, 13, 549, 548, 191, 547, 3,
	21, 19, 541, 540, 536, 535, 2, 17, 9, 534,
	7, 533, 15, 532, 4, 530, 529, 16, 528, 527,
	523, 522, 6, 521, 5, 519, 1, 517, 516, 515,
	14, 8, 47, 238, 513, 512, 511, 510, 509, 508,
	0, 33, 503, 11, 502, 501, 12, 500, 498, 496,
	493, 492, 491, 490, 10, 484, 483, 482, 481, 480,
}
var yyR1 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 4, 4, 5, 6, 7,
	80, 80, 72, 72, 72, 85, 85, 85, 85, 85,
	77, 77, 77, 78, 78, 82, 82, 82, 82, 82,
	82, 82, 83, 83, 83, 83, 83, 83, 83, 84,
	84, 76, 76, 79, 79, 86, 86, 86, 86, 86,
	86, 86, 81, 81, 87, 87, 88, 88, 73, 74,
	74, 75, 8, 8, 8, 9, 9, 9, 10, 11,
	11, 11, 12, 13, 13, 13, 89, 14, 15, 15,
	16, 16, 16, 16, 16, 17, 17, 18, 18, 19,
	19, 19, 22, 22, 20, 20, 20, 23, 23, 24,
	24, 24, 24, 21, 21, 21, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 26, 26, 26, 27, 27,
	28, 28, 28, 28, 29, 29, 30, 30, 31, 31,
	31, 31, 31, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 33, 33, 33, 33, 33, 33, 33,
	37, 37, 37, 42, 38, 38, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 41, 41, 43, 43, 43, 45, 48,
	48, 46, 46, 47, 49, 49, 44, 44, 35, 35,
	35, 35, 50, 50, 51, 51, 52, 52, 53, 53,
	54, 55, 55, 55, 56, 56, 56, 57, 57, 57,
	58, 58, 59, 59, 60, 60, 34, 34, 39, 39,
	40, 40, 61, 61, 62, 63, 63, 64, 64, 65,
	65, 66, 66, 66, 66, 66, 67, 67, 68, 68,
	69, 69, 70, 71,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 12, 3, 7, 7, 8, 7, 3,
	0, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 1, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 0,
	5, 0, 3, 0, 1, 0, 3, 2, 3, 3,
	2, 2, 1, 1, 2, 1, 1, 2, 3, 1,
	3, 7, 1, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 3,
	3, 3, 5, 0, 1, 2, 1, 1, 2, 3,
	2, 3, 2, 2, 2, 1, 3, 1, 1, 3,
	0, 5, 5, 5, 1, 3, 0, 2, 1, 3,
	3, 2, 3, 3, 3, 4, 3, 4, 5, 6,
	3, 4, 2, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 3, 1, 3, 0, 5, 2, 1, 1, 3,
	3, 1, 1, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, 5, 6, 7, 8, 33, -75,
	86, 87, 89, 88, 90, 98, 99, 100, -16, 52,
	53, 54, 55, -14, -89, -14, -14, -14, -14, 91,
	-68, 93, 51, -65, 93, 95, 91, 91, 92, 93,
	91, -71, -71, -71, -3, 17, -17, 18, -15, 29,
	-27, 35, 9, -61, -62, -44, -70, 35, -64, 96,
	92, -70, 35, 91, -70, 35, -63, 96, -70, -63,
	35, -18, -19, 76, -22, 35, -31, -36, -32, 70,
	45, -35, -44, -40, -43, -70, -41, -45, 20, 36,
	37, 38, 25, -42, 74, 75, 49, 96, 28, 81,
	40, -27, 33, 79, -27, 56, 46, 79, 35, 70,
	-70, -71, 35, -71, 94, 35, 20, 67, -70, 9,
	56, -20, -70, 19, 79, 69, 68, -33, 21, 70,
	23, 24, 22, 71, 72, 73, 74, 75, 76, 77,
	78, 46, 47, 48, 41, 42, 43, 44, -31, -36,
	-31, -3, -38, -36, -36, 45, 45, 45, -42, 45,
	-48, -36, -58, 33, 45, -61, 35, -30, 10, -62,
	-36, -70, 45, 20, -69, 97, -66, 89, 87, 32,
	88, 13, 35, 35, 35, -71, -23, -24, -26, 45,
	35, -42, -19, -70, 76, -31, -31, -36, -37, 45,
	-42, 39, 21, 23, 24, -36, -36, 25, 70, -36,
	-36, -36, -36, -36, -36, -36, -36, 128, 128, 56,
	128, -36, 128, -18, 18, -18, -46, -47, 82, -34,
	28, -3, -61, -59, -44, -30, -52, 13, -31, -74,
	-73, 35, 67, -70, -71, -67, 94, -30, 56, -25,
	57, 58, 59, 60, 61, 63, 64, -21, 35, 19,
	-24, 79, -38, -37, -36, -36, 69, 25, -36, 128,
	-18, 128, -49, -47, 84, -31, -60, 67, -39, -40,
	-60, 128, 56, -52, -56, 15, 14, 56, 128, -72,
	-78, -77, -85, -82, -83, 121, 122, 120, 115, 116,
	117, 118, 119, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 113, 114, 35, 35, -50, 11, -24,
	-24, 57, 62, 57, 62, 57, 57, 57, -28, 65,
	95, 66, 35, 128, 35, 128, 69, -36, 128, 85,
	-36, 83, 30, 56, -44, -56, -36, -53, -54, -36,
	-73, -86, -79, 111, -76, 45, -76, -76, -84, 45,
	-84, -84, -84, -76, 45, -84, -76, -71, -51, 12,
	14, 67, 57, 57, 92, 92, 92, -36, -36, 31,
	-40, 56, 56, -55, 26, 27, 70, 25, 32, 124,
	-81, -87, -88, 50, 31, 51, -80, 112, 37, 37,
	37, -52, -31, -38, -31, 45, 45, 45, 7, -36,
	-54, 25, 36, 37, 31, 31, 128, 56, -56, -29,
	-70, -29, -29, -61, 37, -57, 16, 34, 128, 56,
	128, 128, 128, 7, 21, -70, -70, -70,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 86, 86, 86, 86, 86, 72,
	248, 239, 0, 0, 0, 253, 253, 253, 0, 90,
	92, 93, 94, 95, 88, 0, 0, 0, 0, 237,
	0, 0, 249, 0, 0, 240, 0, 235, 0, 235,
	0, 83, 84, 85, 14, 91, 0, 96, 87, 0,
	0, 128, 0, 19, 232, 0, 196, 252, 0, 0,
	0, 253, 252, 0, 253, 0, 0, 0, 0, 0,
	82, 0, 97, 99, 104, 252, 102, 103, 138, 0,
	0, 166, 167, 168, 0, 196, 0, 182, 0, 198,
	199, 200, 201, 231, 185, 186, 187, 183, 184, 189,
	89, 220, 0, 0, 136, 0, 0, 0, 0, 0,
	250, 74, 0, 77, 0, 79, 236, 0, 253, 0,
	0, 100, 105, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 153, 154, 155, 156, 157, 158, 159, 141, 0,
	0, 0, 0, 164, 177, 0, 0, 0, 152, 0,
	0, 190, 0, 0, 0, 136, 129, 206, 0, 233,
	234, 197, 0, 238, 0, 0, 253, 246, 241, 242,
	243, 244, 245, 78, 80, 81, 136, 107, 113, 0,
	125, 127, 98, 106, 101, 139, 140, 143, 144, 0,
	161, 162, 0, 0, 0, 146, 0, 150, 0, 169,
	170, 171, 172, 173, 174, 175, 176, 142, 163, 0,
	230, 164, 178, 0, 0, 0, 194, 191, 0, 224,
	0, 227, 224, 0, 222, 206, 214, 0, 137, 0,
	69, 0, 0, 251, 75, 0, 247, 202, 0, 0,
	116, 117, 0, 0, 0, 0, 0, 130, 114, 0,
	0, 0, 0, 145, 147, 0, 0, 151, 165, 179,
	0, 181, 0, 192, 0, 0, 15, 0, 226, 228,
	16, 221, 0, 214, 18, 0, 0, 0, 71, 55,
	53, 23, 24, 51, 34, 51, 51, 32, 25, 26,
	27, 28, 29, 35, 36, 37, 38, 39, 40, 41,
	49, 49, 49, 49, 49, 253, 76, 204, 0, 108,
	111, 118, 0, 120, 0, 122, 123, 124, 109, 0,
	0, 0, 115, 110, 126, 160, 0, 148, 180, 188,
	195, 0, 0, 0, 223, 17, 215, 207, 208, 211,
	70, 68, 20, 54, 33, 0, 30, 31, 42, 0,
	43, 44, 45, 46, 0, 47, 48, 73, 206, 0,
	0, 0, 119, 121, 0, 0, 0, 149, 193, 0,
	229, 0, 0, 210, 212, 213, 0, 57, 0, 60,
	61, 62, 63, 0, 65, 66, 22, 21, 0, 0,
	0, 214, 205, 203, 112, 0, 0, 0, 0, 216,
	209, 56, 58, 59, 64, 67, 52, 0, 217, 0,
	134, 0, 0, 225, 0, 13, 0, 0, 131, 0,
	132, 133, 50, 218, 0, 135, 0, 219,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 78, 71, 3,
	45, 128, 76, 74, 56, 75, 79, 77, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	47, 46, 48, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 73, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 72, 3, 49,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 50, 51, 52, 53, 54, 55, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 70, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:181
		{
			SetParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:187
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:203
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(AST_WHERE, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(AST_HAVING, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:207
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 15:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:213
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: yyDollar[5].columns, Rows: yyDollar[6].insRows, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 16:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:217
		{
			cols := make(Columns, 0, len(yyDollar[6].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[6].updateExprs))
			for _, col := range yyDollar[6].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[7].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:229
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(AST_WHERE, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 18:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:235
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(AST_WHERE, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:241
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:246
		{
			yyVAL.str = ""
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:250
		{
			yyVAL.str = AST_ZEROFILL
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:255
		{
			yyVAL.str = yyDollar[1].str
			if yyDollar[2].str != "" {
				yyVAL.str += " " + yyDollar[2].str
			}
			if yyDollar[3].str != "" {
				yyVAL.str += " " + yyDollar[3].str
			}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:269
		{
			yyVAL.str = AST_DATE
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:273
		{
			yyVAL.str = AST_TIME
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:277
		{
			yyVAL.str = AST_TIMESTAMP
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:281
		{
			yyVAL.str = AST_DATETIME
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:285
		{
			yyVAL.str = AST_YEAR
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:291
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_CHAR
			} else {
				yyVAL.str = AST_CHAR + yyDollar[2].str
			}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:299
		{
			if yyDollar[2].str == "" {
				yyVAL.str = AST_VARCHAR
			} else {
				yyVAL.str = AST_VARCHAR + yyDollar[2].str
			}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:307
		{
			yyVAL.str = AST_TEXT
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:313
		{
			yyVAL.str = yyDollar[1].str + yyDollar[2].str
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:317
		{
			yyVAL.str = yyDollar[1].str
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:323
		{
			yyVAL.str = AST_BIT
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:327
		{
			yyVAL.str = AST_TINYINT
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:331
		{
			yyVAL.str = AST_SMALLINT
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:335
		{
			yyVAL.str = AST_MEDIUMINT
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:339
		{
			yyVAL.str = AST_INT
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:343
		{
			yyVAL.str = AST_INTEGER
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:347
		{
			yyVAL.str = AST_BIGINT
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:353
		{
			yyVAL.str = AST_REAL + yyDollar[2].str
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:357
		{
			yyVAL.str = AST_DOUBLE + yyDollar[2].str
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:361
		{
			yyVAL.str = AST_FLOAT + yyDollar[2].str
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:365
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:369
		{
			yyVAL.str = AST_DECIMAL + yyDollar[2].str
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:373
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:377
		{
			yyVAL.str = AST_NUMERIC + yyDollar[2].str
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:382
		{
			yyVAL.str = ""
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:386
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ", " + string(yyDollar[4].bytes) + ")"
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = ""
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = "(" + string(yyDollar[2].bytes) + ")"
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:400
		{
			yyVAL.str = ""
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:404
		{
			yyVAL.str = AST_UNSIGNED
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:409
		{
			yyVAL.columnAtts = ColumnAtts{}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:413
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_NOT_NULL)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:419
		{
			node := StrVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:424
		{
			node := NumVal(yyDollar[3].bytes)
			yyVAL.columnAtts = append(yyVAL.columnAtts, "default "+String(node))
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:429
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, AST_AUTO_INCREMENT)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:433
		{
			yyVAL.columnAtts = append(yyVAL.columnAtts, yyDollar[2].str)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:439
		{
			yyVAL.str = AST_PRIMARY_KEY
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:443
		{
			yyVAL.str = AST_UNIQUE_KEY
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:457
		{
			yyVAL.columnDefinition = &ColumnDefinition{ColName: string(yyDollar[1].bytes), ColType: yyDollar[2].str, ColumnAtts: yyDollar[3].columnAtts}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:463
		{
			yyVAL.columnDefinitions = ColumnDefinitions{yyDollar[1].columnDefinition}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:467
		{
			yyVAL.columnDefinitions = append(yyVAL.columnDefinitions, yyDollar[3].columnDefinition)
		}
	case 71:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:473
		{
			yyVAL.statement = &CreateTable{Name: yyDollar[4].bytes, ColumnDefinitions: yyDollar[6].columnDefinitions}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:479
		{
			yyVAL.statement = yyDollar[1].statement
		}
	case 73:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:483
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[7].bytes, NewName: yyDollar[7].bytes}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:488
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyDollar[3].bytes}
		}
	case 75:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:494
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[4].bytes, NewName: yyDollar[4].bytes}
		}
	case 76:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:498
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[4].bytes, NewName: yyDollar[7].bytes}
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:503
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 78:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:509
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyDollar[3].bytes, NewName: yyDollar[5].bytes}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:515
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:519
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[5].bytes, NewName: yyDollar[5].bytes}
		}
	case 81:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:524
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyDollar[4].bytes}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:530
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyDollar[3].bytes, NewName: yyDollar[3].bytes}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:536
		{
			yyVAL.statement = &Other{}
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:540
		{
			yyVAL.statement = &Other{}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:544
		{
			yyVAL.statement = &Other{}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:549
		{
			SetAllowComments(yylex, true)
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:553
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			SetAllowComments(yylex, false)
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:559
		{
			yyVAL.bytes2 = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:563
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:569
		{
			yyVAL.str = AST_UNION
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:573
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:577
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:581
		{
			yyVAL.str = AST_EXCEPT
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:585
		{
			yyVAL.str = AST_INTERSECT
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:590
		{
			yyVAL.str = ""
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:594
		{
			yyVAL.str = AST_DISTINCT
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:600
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:604
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:610
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:614
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].bytes}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:618
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].bytes}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:624
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:628
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:633
		{
			yyVAL.bytes = nil
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:637
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:641
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:647
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:651
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:657
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].bytes, Hints: yyDollar[3].indexHints}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:661
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyDollar[2].tableExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:665
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 112:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:669
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 113:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:674
		{
			yyVAL.bytes = nil
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:678
		{
			yyVAL.bytes = yyDollar[1].bytes
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:682
		{
			yyVAL.bytes = yyDollar[2].bytes
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:688
		{
			yyVAL.str = AST_JOIN
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:692
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 118:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:696
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:700
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:704
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:708
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:712
		{
			yyVAL.str = AST_JOIN
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:716
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:720
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:726
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].bytes}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:730
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:734
		{
			yyVAL.smTableExpr = yyDollar[1].subquery
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].bytes}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:744
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:749
		{
			yyVAL.indexHints = nil
		}
	case 131:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:753
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyDollar[4].bytes2}
		}
	case 132:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:757
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyDollar[4].bytes2}
		}
	case 133:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:761
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyDollar[4].bytes2}
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:767
		{
			yyVAL.bytes2 = [][]byte{yyDollar[1].bytes}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:771
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[3].bytes)
		}
	case 136:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:776
		{
			yyVAL.boolExpr = nil
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:780
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:787
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:791
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:795
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:805
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:809
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_IN, Right: yyDollar[3].colTuple}
		}
	case 145:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:813
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_IN, Right: yyDollar[4].colTuple}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:817
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_LIKE, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:821
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: AST_NOT_LIKE, Right: yyDollar[4].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:825
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_BETWEEN, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:829
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: AST_NOT_BETWEEN, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:833
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyDollar[1].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:837
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyDollar[1].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:841
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:847
		{
			yyVAL.str = AST_EQ
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:851
		{
			yyVAL.str = AST_LT
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:855
		{
			yyVAL.str = AST_GT
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:859
		{
			yyVAL.str = AST_LE
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:863
		{
			yyVAL.str = AST_GE
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:867
		{
			yyVAL.str = AST_NE
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:871
		{
			yyVAL.str = AST_NSE
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:877
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:881
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:885
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:891
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 165:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:901
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:907
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:911
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:915
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITAND, Right: yyDollar[3].valExpr}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITOR, Right: yyDollar[3].valExpr}
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:927
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_BITXOR, Right: yyDollar[3].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:931
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_PLUS, Right: yyDollar[3].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:935
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MINUS, Right: yyDollar[3].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:939
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MULT, Right: yyDollar[3].valExpr}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_DIV, Right: yyDollar[3].valExpr}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: AST_MOD, Right: yyDollar[3].valExpr}
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:951
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				switch yyDollar[1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyDollar[1].byt, Expr: yyDollar[2].valExpr}
			}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:966
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes}
		}
	case 179:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:970
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 180:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:974
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 181:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:978
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].bytes, Exprs: yyDollar[3].selectExprs}
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:982
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:988
		{
			yyVAL.bytes = IF_BYTES
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:992
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:998
		{
			yyVAL.byt = AST_UPLUS
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.byt = AST_UMINUS
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.byt = AST_TILDA
		}
	case 188:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1017
		{
			yyVAL.valExpr = nil
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1027
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1031
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.valExpr = nil
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1052
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].bytes}
		}
	case 197:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].bytes, Name: yyDollar[3].bytes}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1062
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1066
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.valExprs = nil
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.boolExpr = nil
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 206:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.orderBy = nil
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1122
		{
			yyVAL.str = AST_ASC
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.str = AST_ASC
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.str = AST_DESC
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1135
		{
			yyVAL.limit = nil
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1139
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 216:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.str = ""
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 219:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1156
		{
			if !bytes.Equal(yyDollar[3].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyDollar[4].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1169
		{
			yyVAL.columns = nil
		}
	case 221:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1173
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1179
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1183
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1188
		{
			yyVAL.updateExprs = nil
		}
	case 225:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1192
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 226:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1198
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1202
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1208
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1212
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1218
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1222
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1228
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1232
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 234:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1238
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1243
		{
			yyVAL.empty = struct{}{}
		}
	case 236:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1245
		{
			yyVAL.empty = struct{}{}
		}
	case 237:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1248
		{
			yyVAL.empty = struct{}{}
		}
	case 238:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1250
		{
			yyVAL.empty = struct{}{}
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1253
		{
			yyVAL.empty = struct{}{}
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1255
		{
			yyVAL.empty = struct{}{}
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1259
		{
			yyVAL.empty = struct{}{}
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1261
		{
			yyVAL.empty = struct{}{}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1263
		{
			yyVAL.empty = struct{}{}
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1265
		{
			yyVAL.empty = struct{}{}
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1267
		{
			yyVAL.empty = struct{}{}
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1270
		{
			yyVAL.empty = struct{}{}
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1272
		{
			yyVAL.empty = struct{}{}
		}
	case 248:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1275
		{
			yyVAL.empty = struct{}{}
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1277
		{
			yyVAL.empty = struct{}{}
		}
	case 250:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 251:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1282
		{
			yyVAL.empty = struct{}{}
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1286
		{
			yyVAL.bytes = bytes.ToLower(yyDollar[1].bytes)
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1291
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
