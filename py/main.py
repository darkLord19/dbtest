import psycopg2

conn_str = "postgresql://postgres:postgres@localhost:5432/postgres"
conn = psycopg2.connect(conn_str)

with conn:
    with conn.cursor() as curs:
        query = "INSERT INTO test(col1, col2, col3) VALUES ('hello', 1, 1111)"
        curs.execute(query)
        query = "SELECT * FROM test"
        curs.execute(query)
        rows = curs.fetchall()
        for row in rows:
            print(row)

conn.close()
