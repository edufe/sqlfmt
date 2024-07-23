package sqlfmt

// Derived from PostgreSQL -- ./src/include/parser/kwlist.h

var keywords map[string]int

func init() {
	keywords = make(map[string]int)
	/* name, value, category */
	keywords["abort"] = ABORT_P
	keywords["absolute"] = ABSOLUTE_P
	keywords["access"] = ACCESS
	keywords["action"] = ACTION
	keywords["add"] = ADD_P
	keywords["admin"] = ADMIN
	keywords["after"] = AFTER
	keywords["aggregate"] = AGGREGATE
	keywords["all"] = ALL
	keywords["also"] = ALSO
	keywords["alter"] = ALTER
	keywords["always"] = ALWAYS
	keywords["analyse"] = ANALYSE /* British spelling */
	keywords["analyze"] = ANALYZE
	keywords["and"] = AND
	keywords["any"] = ANY
	keywords["array"] = ARRAY
	keywords["as"] = AS
	keywords["asc"] = ASC
	keywords["assertion"] = ASSERTION
	keywords["assignment"] = ASSIGNMENT
	keywords["asymmetric"] = ASYMMETRIC
	keywords["at"] = AT
	keywords["attribute"] = ATTRIBUTE
	keywords["authorization"] = AUTHORIZATION
	keywords["backward"] = BACKWARD
	keywords["before"] = BEFORE
	keywords["begin"] = BEGIN_P
	keywords["between"] = BETWEEN
	keywords["bigint"] = BIGINT
	keywords["binary"] = BINARY
	keywords["bit"] = BIT
	keywords["boolean"] = BOOLEAN_P
	keywords["both"] = BOTH
	keywords["by"] = BY
	keywords["cache"] = CACHE
	keywords["called"] = CALLED
	keywords["cascade"] = CASCADE
	keywords["cascaded"] = CASCADED
	keywords["case"] = CASE
	keywords["cast"] = CAST
	keywords["catalog"] = CATALOG_P
	keywords["chain"] = CHAIN
	keywords["char"] = CHAR_P
	keywords["character"] = CHARACTER
	keywords["characteristics"] = CHARACTERISTICS
	keywords["check"] = CHECK
	keywords["checkpoint"] = CHECKPOINT
	keywords["class"] = CLASS
	keywords["close"] = CLOSE
	keywords["cluster"] = CLUSTER
	keywords["coalesce"] = COALESCE
	keywords["collate"] = COLLATE
	keywords["collation"] = COLLATION
	keywords["column"] = COLUMN
	keywords["comment"] = COMMENT
	keywords["comments"] = COMMENTS
	keywords["commit"] = COMMIT
	keywords["committed"] = COMMITTED
	keywords["concurrently"] = CONCURRENTLY
	keywords["configuration"] = CONFIGURATION
	keywords["conflict"] = CONFLICT
	keywords["connection"] = CONNECTION
	keywords["constraint"] = CONSTRAINT
	keywords["constraints"] = CONSTRAINTS
	keywords["content"] = CONTENT_P
	keywords["continue"] = CONTINUE_P
	keywords["conversion"] = CONVERSION_P
	keywords["copy"] = COPY
	keywords["cost"] = COST
	keywords["create"] = CREATE
	keywords["cross"] = CROSS
	keywords["csv"] = CSV
	keywords["cube"] = CUBE
	keywords["current"] = CURRENT_P
	keywords["current_catalog"] = CURRENT_CATALOG
	keywords["current_date"] = CURRENT_DATE
	keywords["current_role"] = CURRENT_ROLE
	keywords["current_schema"] = CURRENT_SCHEMA
	keywords["current_time"] = CURRENT_TIME
	keywords["current_timestamp"] = CURRENT_TIMESTAMP
	keywords["current_user"] = CURRENT_USER
	keywords["cursor"] = CURSOR
	keywords["cycle"] = CYCLE
	keywords["data"] = DATA_P
	keywords["database"] = DATABASE
	keywords["day"] = DAY_P
	keywords["deallocate"] = DEALLOCATE
	keywords["dec"] = DEC
	keywords["decimal"] = DECIMAL_P
	keywords["declare"] = DECLARE
	keywords["default"] = DEFAULT
	keywords["defaults"] = DEFAULTS
	keywords["deferrable"] = DEFERRABLE
	keywords["deferred"] = DEFERRED
	keywords["definer"] = DEFINER
	keywords["delete"] = DELETE_P
	keywords["delimiter"] = DELIMITER
	keywords["delimiters"] = DELIMITERS
	keywords["desc"] = DESC
	keywords["dictionary"] = DICTIONARY
	keywords["disable"] = DISABLE_P
	keywords["discard"] = DISCARD
	keywords["distinct"] = DISTINCT
	keywords["do"] = DO
	keywords["document"] = DOCUMENT_P
	keywords["domain"] = DOMAIN_P
	keywords["double"] = DOUBLE_P
	keywords["drop"] = DROP
	keywords["each"] = EACH
	keywords["else"] = ELSE
	keywords["enable"] = ENABLE_P
	keywords["encoding"] = ENCODING
	keywords["encrypted"] = ENCRYPTED
	keywords["end"] = END_P
	keywords["enum"] = ENUM_P
	keywords["escape"] = ESCAPE
	keywords["event"] = EVENT
	keywords["except"] = EXCEPT
	keywords["exclude"] = EXCLUDE
	keywords["excluding"] = EXCLUDING
	keywords["exclusive"] = EXCLUSIVE
	keywords["execute"] = EXECUTE
	keywords["exists"] = EXISTS
	keywords["explain"] = EXPLAIN
	keywords["extension"] = EXTENSION
	keywords["external"] = EXTERNAL
	keywords["extract"] = EXTRACT
	keywords["false"] = FALSE_P
	keywords["family"] = FAMILY
	keywords["fetch"] = FETCH
	keywords["filter"] = FILTER
	keywords["first"] = FIRST_P
	keywords["float"] = FLOAT_P
	keywords["following"] = FOLLOWING
	keywords["for"] = FOR
	keywords["force"] = FORCE
	keywords["foreign"] = FOREIGN
	keywords["forward"] = FORWARD
	keywords["freeze"] = FREEZE
	keywords["from"] = FROM
	keywords["full"] = FULL
	keywords["function"] = FUNCTION
	keywords["functions"] = FUNCTIONS
	keywords["global"] = GLOBAL
	keywords["grant"] = GRANT
	keywords["granted"] = GRANTED
	keywords["greatest"] = GREATEST
	keywords["group"] = GROUP_P
	keywords["grouping"] = GROUPING
	keywords["handler"] = HANDLER
	keywords["having"] = HAVING
	keywords["header"] = HEADER_P
	keywords["hold"] = HOLD
	keywords["hour"] = HOUR_P
	keywords["identity"] = IDENTITY_P
	keywords["if"] = IF_P
	keywords["ilike"] = ILIKE
	keywords["immediate"] = IMMEDIATE
	keywords["immutable"] = IMMUTABLE
	keywords["implicit"] = IMPLICIT_P
	keywords["import"] = IMPORT_P
	keywords["in"] = IN_P
	keywords["including"] = INCLUDING
	keywords["increment"] = INCREMENT
	keywords["index"] = INDEX
	keywords["indexes"] = INDEXES
	keywords["inherit"] = INHERIT
	keywords["inherits"] = INHERITS
	keywords["initially"] = INITIALLY
	keywords["inline"] = INLINE_P
	keywords["inner"] = INNER_P
	keywords["inout"] = INOUT
	keywords["input"] = INPUT_P
	keywords["insensitive"] = INSENSITIVE
	keywords["insert"] = INSERT
	keywords["instead"] = INSTEAD
	keywords["int"] = INT_P
	keywords["integer"] = INTEGER
	keywords["intersect"] = INTERSECT
	keywords["interval"] = INTERVAL
	keywords["into"] = INTO
	keywords["invoker"] = INVOKER
	keywords["is"] = IS
	keywords["isnull"] = ISNULL
	keywords["isolation"] = ISOLATION
	keywords["join"] = JOIN
	keywords["key"] = KEY
	keywords["label"] = LABEL
	keywords["language"] = LANGUAGE
	keywords["large"] = LARGE_P
	keywords["last"] = LAST_P
	keywords["lateral"] = LATERAL_P
	keywords["leading"] = LEADING
	keywords["leakproof"] = LEAKPROOF
	keywords["least"] = LEAST
	keywords["left"] = LEFT
	keywords["level"] = LEVEL
	keywords["like"] = LIKE
	keywords["limit"] = LIMIT
	keywords["listen"] = LISTEN
	keywords["load"] = LOAD
	keywords["local"] = LOCAL
	keywords["localtime"] = LOCALTIME
	keywords["localtimestamp"] = LOCALTIMESTAMP
	keywords["location"] = LOCATION
	keywords["lock"] = LOCK_P
	keywords["locked"] = LOCKED
	keywords["logged"] = LOGGED
	keywords["mapping"] = MAPPING
	keywords["match"] = MATCH
	keywords["materialized"] = MATERIALIZED
	keywords["maxvalue"] = MAXVALUE
	keywords["minute"] = MINUTE_P
	keywords["minvalue"] = MINVALUE
	keywords["mode"] = MODE
	keywords["month"] = MONTH_P
	keywords["move"] = MOVE
	keywords["name"] = NAME_P
	keywords["names"] = NAMES
	keywords["national"] = NATIONAL
	keywords["natural"] = NATURAL
	keywords["nchar"] = NCHAR
	keywords["next"] = NEXT
	keywords["no"] = NO
	keywords["none"] = NONE
	keywords["not"] = NOT
	keywords["nothing"] = NOTHING
	keywords["notify"] = NOTIFY
	keywords["notnull"] = NOTNULL
	keywords["nowait"] = NOWAIT
	keywords["null"] = NULL_P
	keywords["nullif"] = NULLIF
	keywords["nulls"] = NULLS_P
	keywords["numeric"] = NUMERIC
	keywords["object"] = OBJECT_P
	keywords["of"] = OF
	keywords["off"] = OFF
	keywords["offset"] = OFFSET
	keywords["oids"] = OIDS
	keywords["on"] = ON
	keywords["only"] = ONLY
	keywords["operator"] = OPERATOR
	keywords["option"] = OPTION
	keywords["options"] = OPTIONS
	keywords["or"] = OR
	keywords["order"] = ORDER
	keywords["ordinality"] = ORDINALITY
	keywords["out"] = OUT_P
	keywords["outer"] = OUTER_P
	keywords["over"] = OVER
	keywords["overlaps"] = OVERLAPS
	keywords["overlay"] = OVERLAY
	keywords["owned"] = OWNED
	keywords["owner"] = OWNER
	keywords["parser"] = PARSER
	keywords["partial"] = PARTIAL
	keywords["partition"] = PARTITION
	keywords["passing"] = PASSING
	keywords["password"] = PASSWORD
	keywords["placing"] = PLACING
	keywords["plans"] = PLANS
	keywords["policy"] = POLICY
	keywords["position"] = POSITION
	keywords["preceding"] = PRECEDING
	keywords["precision"] = PRECISION
	keywords["prepare"] = PREPARE
	keywords["prepared"] = PREPARED
	keywords["preserve"] = PRESERVE
	keywords["primary"] = PRIMARY
	keywords["prior"] = PRIOR
	keywords["privileges"] = PRIVILEGES
	keywords["procedural"] = PROCEDURAL
	keywords["procedure"] = PROCEDURE
	keywords["program"] = PROGRAM
	keywords["quote"] = QUOTE
	keywords["range"] = RANGE
	keywords["read"] = READ
	keywords["real"] = REAL
	keywords["reassign"] = REASSIGN
	keywords["recheck"] = RECHECK
	keywords["recursive"] = RECURSIVE
	keywords["ref"] = REF
	keywords["references"] = REFERENCES
	keywords["refresh"] = REFRESH
	keywords["reindex"] = REINDEX
	keywords["relative"] = RELATIVE_P
	keywords["release"] = RELEASE
	keywords["rename"] = RENAME
	keywords["repeatable"] = REPEATABLE
	keywords["replace"] = REPLACE
	keywords["replica"] = REPLICA
	keywords["reset"] = RESET
	keywords["restart"] = RESTART
	keywords["restrict"] = RESTRICT
	keywords["returning"] = RETURNING
	keywords["returns"] = RETURNS
	keywords["revoke"] = REVOKE
	keywords["right"] = RIGHT
	keywords["role"] = ROLE
	keywords["rollback"] = ROLLBACK
	keywords["rollup"] = ROLLUP
	keywords["row"] = ROW
	keywords["rows"] = ROWS
	keywords["rule"] = RULE
	keywords["savepoint"] = SAVEPOINT
	keywords["schema"] = SCHEMA
	keywords["scroll"] = SCROLL
	keywords["search"] = SEARCH
	keywords["second"] = SECOND_P
	keywords["security"] = SECURITY
	keywords["select"] = SELECT
	keywords["sequence"] = SEQUENCE
	keywords["sequences"] = SEQUENCES
	keywords["serializable"] = SERIALIZABLE
	keywords["server"] = SERVER
	keywords["session"] = SESSION
	keywords["session_user"] = SESSION_USER
	keywords["set"] = SET
	keywords["setof"] = SETOF
	keywords["sets"] = SETS
	keywords["share"] = SHARE
	keywords["show"] = SHOW
	keywords["similar"] = SIMILAR
	keywords["simple"] = SIMPLE
	keywords["skip"] = SKIP
	keywords["smallint"] = SMALLINT
	keywords["snapshot"] = SNAPSHOT
	keywords["some"] = SOME
	keywords["sql"] = SQL_P
	keywords["stable"] = STABLE
	keywords["standalone"] = STANDALONE_P
	keywords["start"] = START
	keywords["statement"] = STATEMENT
	keywords["statistics"] = STATISTICS
	keywords["stdin"] = STDIN
	keywords["stdout"] = STDOUT
	keywords["storage"] = STORAGE
	keywords["strict"] = STRICT_P
	keywords["strip"] = STRIP_P
	keywords["substr"] = SUBSTRING
	keywords["substring"] = SUBSTRING
	keywords["symmetric"] = SYMMETRIC
	keywords["sysid"] = SYSID
	keywords["system"] = SYSTEM_P
	keywords["table"] = TABLE
	keywords["tables"] = TABLES
	keywords["tablesample"] = TABLESAMPLE
	keywords["tablespace"] = TABLESPACE
	keywords["temp"] = TEMP
	keywords["template"] = TEMPLATE
	keywords["temporary"] = TEMPORARY
	keywords["text"] = TEXT_P
	keywords["then"] = THEN
	keywords["time"] = TIME
	keywords["timestamp"] = TIMESTAMP
	keywords["to"] = TO
	keywords["trailing"] = TRAILING
	keywords["transaction"] = TRANSACTION
	keywords["transform"] = TRANSFORM
	keywords["treat"] = TREAT
	keywords["trigger"] = TRIGGER
	keywords["trim"] = TRIM
	keywords["true"] = TRUE_P
	keywords["truncate"] = TRUNCATE
	keywords["trusted"] = TRUSTED
	keywords["type"] = TYPE_P
	keywords["types"] = TYPES_P
	keywords["unbounded"] = UNBOUNDED
	keywords["uncommitted"] = UNCOMMITTED
	keywords["unencrypted"] = UNENCRYPTED
	keywords["union"] = UNION
	keywords["unique"] = UNIQUE
	keywords["unknown"] = UNKNOWN
	keywords["unlisten"] = UNLISTEN
	keywords["unlogged"] = UNLOGGED
	keywords["until"] = UNTIL
	keywords["update"] = UPDATE
	keywords["user"] = USER
	keywords["using"] = USING
	keywords["vacuum"] = VACUUM
	keywords["valid"] = VALID
	keywords["validate"] = VALIDATE
	keywords["validator"] = VALIDATOR
	keywords["value"] = VALUE_P
	keywords["values"] = VALUES
	keywords["varchar"] = VARCHAR
	keywords["variadic"] = VARIADIC
	keywords["varying"] = VARYING
	keywords["verbose"] = VERBOSE
	keywords["version"] = VERSION_P
	keywords["view"] = VIEW
	keywords["views"] = VIEWS
	keywords["volatile"] = VOLATILE
	keywords["when"] = WHEN
	keywords["where"] = WHERE
	keywords["whitespace"] = WHITESPACE_P
	keywords["window"] = WINDOW
	keywords["with"] = WITH
	keywords["within"] = WITHIN
	keywords["without"] = WITHOUT
	keywords["work"] = WORK
	keywords["wrapper"] = WRAPPER
	keywords["write"] = WRITE
	keywords["xml"] = XML_P
	keywords["xmlattributes"] = XMLATTRIBUTES
	keywords["xmlconcat"] = XMLCONCAT
	keywords["xmlelement"] = XMLELEMENT
	keywords["xmlexists"] = XMLEXISTS
	keywords["xmlforest"] = XMLFOREST
	keywords["xmlparse"] = XMLPARSE
	keywords["xmlpi"] = XMLPI
	keywords["xmlroot"] = XMLROOT
	keywords["xmlserialize"] = XMLSERIALIZE
	keywords["year"] = YEAR_P
	keywords["yes"] = YES_P
	keywords["zone"] = ZONE
}
