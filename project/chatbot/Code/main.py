from flask import Flask, request, jsonify
from chat_bot import DocumentChatbot
import sys
#import google_chat
import connection

app = Flask(__name__)

chatbot = DocumentChatbot()
def main():
    #test.test_bot()
    print("chatbot runing...", "*"*50 ,file=sys.stderr)
    data = connection.retrieve_data()
    if data:
        connection.save_to_txt(data)
    else:
        print("No data retrieved.")
    run_flask_app() 
    #chat_bot.run()
    #google_chat.main()


@app.route('/webhook', methods=['POST'])
def webhook():
    data = request.get_json(silent=True)
    if not data:
        print("No JSON received", file=sys.stderr)
        return jsonify({"error": "No JSON received"}), 400
    print("Received message:", data, file=sys.stderr)
    message_text = chatbot.process_query(data['message'])
    return jsonify({"text": f"{message_text}"})

def run_flask_app():
    # Run the Flask app
    app.run(debug=True,host='0.0.0.0' ,port=5000)

if __name__ == '__main__':
    main()

