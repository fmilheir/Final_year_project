import psycopg2
import os
import sys

def retrieve_data():
    
    # PostgreSQL connection parameters
    conn_params = {
        "host": "db",
        "database": os.getenv("DB_NAME"),
        "user": os.getenv("DB_USER"),
        "password": os.getenv("DB_PASSWORD"),
        "port": "5432"
    }
    print("my dbname is" ,os.getenv("DB_NAME") ,file=sys.stderr)
    print("my dbuser is" ,os.getenv("DB_USER") ,file=sys.stderr)
    print("my dbpassword is" ,os.getenv("DB_PASSWORD") ,file=sys.stderr)

    # Connect to the PostgreSQL database
    try:
        conn = psycopg2.connect(**conn_params)
        cursor = conn.cursor()
    except psycopg2.Error as e:
        print("Error: Unable to connect to the database.")
        print(e)
        return None

    # Retrieve data from the database
    try:
        cursor.execute("SELECT * FROM incidents;")
        data = cursor.fetchall()
    except psycopg2.Error as e:
        print("Error: Unable to execute query.")
        print(e)
        cursor.close()
        conn.close()
        return None

    cursor.close()
    conn.close()

    return data

def save_to_txt(data):
    # Check if the PDF directory exists, create it if it doesn't
    pdf_dir = "pdf"
    if not os.path.exists(pdf_dir):
        os.makedirs(pdf_dir)

    # Save data to a text file in the PDF directory
    file_path = os.path.join(pdf_dir, "data.txt")
    with open(file_path, "w") as f:
        for row in data:
            f.write(str(row) + "\n")

    print("Data saved to:", file_path)

