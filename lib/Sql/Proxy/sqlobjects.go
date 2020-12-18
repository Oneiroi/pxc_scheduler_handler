package Proxy

/*
Class dealing with ALL SQL commands.
No SQL should be hardcoded in the any other class
 */


const(
	/*
	Table used as backup / point in time recovery before scheduler start to take control
	 */
	Ddl_Create_mysal_server_original = "CREATE TABLE mysql_servers_original (" +
		"hostgroup_id INT CHECK (hostgroup_id>=0) NOT NULL DEFAULT 0, " +
		"hostname VARCHAR NOT NULL, " +
		"port INT CHECK (port >= 0 AND port <= 65535) NOT NULL DEFAULT 3306, " +
		"gtid_port INT CHECK ((gtid_port <> port OR gtid_port=0) AND gtid_port >= 0 AND gtid_port <= 65535) NOT NULL DEFAULT 0," +
		"status VARCHAR CHECK (UPPER(status) IN ('ONLINE','SHUNNED','OFFLINE_SOFT', 'OFFLINE_HARD')) NOT NULL DEFAULT 'ONLINE', " +
		"weight INT CHECK (weight >= 0 AND weight <=10000000) NOT NULL DEFAULT 1, " +
		"compression INT CHECK (compression IN(0,1)) NOT NULL DEFAULT 0, " +
		"max_connections INT CHECK (max_connections >=0) NOT NULL DEFAULT 1000, " +
		"max_replication_lag INT CHECK (max_replication_lag >= 0 AND max_replication_lag <= 126144000) NOT NULL DEFAULT 0, " +
		"use_ssl INT CHECK (use_ssl IN(0,1)) NOT NULL DEFAULT 0, " +
		"max_latency_ms INT UNSIGNED CHECK (max_latency_ms>=0) NOT NULL DEFAULT 0, " +
		"comment VARCHAR NOT NULL DEFAULT '', " +
		"PRIMARY KEY (hostgroup_id, hostname, port))  "
	/*
	This table is the working area for the scheduler to deal with the nodes while processing
	 */
	Ddl_create_mysql_servers_scheduler = "CREATE TABLE mysql_servers_scheduler (" +
		"hostgroup_id INT CHECK (hostgroup_id>=0) NOT NULL DEFAULT 0, " +
		"hostname VARCHAR NOT NULL,    port INT CHECK (port >= 0 AND port <= 65535) NOT NULL DEFAULT 3306, " +
		"status VARCHAR CHECK (UPPER(status) IN ('ONLINE','SHUNNED','OFFLINE_SOFT', 'OFFLINE_HARD')) NOT NULL DEFAULT 'ONLINE',  " +
		"retry_up INT DEFAULT 0, " +
		"retry_down INT DEFAULT 0, " +
		"previous_status VARCHAR,  " +
		"backup_hg INT, " +
		"PRIMARY KEY (hostgroup_id, hostname, port))"

	/*
	Configuration table for pxc
	 */
	Ddl_create_pxc_clusters = "CREATE TABLE pxc_clusters (cluster_id INTEGER PRIMARY KEY AUTOINCREMENT," +
		"hg_w INT," +
		"hg_r INT," +
		"bck_hg_w INT," +
		"bck_hg_r INT," +
		"single_writer INT DEFAULT 1," +
		"max_writers INT DEFAULT 1," +
		"writer_is_also_reader INT DEFAULT 1, " +
		"retry_up INT DEFAULT 0, " +
		"retry_down INT DEFAULT 0)"

	/*
	cleanup of tables
	 */
	Ddl_drop_mysql_server_original = "DROP TABLE IF EXISTS mysql_servers_original"
	Ddl_drop_mysql_server_scheduler = "DROP TABLE IF EXISTS mysql_servers_scheduler"
	Ddl_drop_pxc_cluster = "DROP TABLE IF EXISTS pxc_clusters"

	Ddl_Truncate_mysql_nodes = "truncate table from mysql_servers"

	/*
	Information retrieval
 	*/


	Dml_Select_mysql_nodes = "select hostgroup_id, hostname,port, status,weight, max_connections, max_replication_lag  from mysql_servers where hostgroup_id = ?"



)