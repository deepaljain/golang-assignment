liquibase-up:
	liquibase --changeLogFile=./migrations/changelog.xml --url="jdbc:postgresql://localhost:5432/user_lb_db" --username=deepaljain --password=postgres update

liquibase-rollback:
	liquibase --changeLogFile=./migrations/changelog.xml --url="jdbc:postgresql://localhost:5432/user_lb_db" --username=deepaljain --password=postgres rollbackCount 1