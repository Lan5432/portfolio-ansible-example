from flask import Flask, render_template, request

app = Flask(__name__)

@app.route("/")
def hello():
    return render_template("index.html", request=request)

if __name__ == "__main__":
    app.run(host='0.0.0.0')
