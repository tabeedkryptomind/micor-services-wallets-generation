from flask import Flask
from flask import jsonify,request
from wallets import *

addr_index = 0
app = Flask(__name__)


def increment():
    global addr_index
    addr_index = addr_index+1
@app.route('/api/btc/create-adderess/', methods=['GET'])
def get_adders():
    pvk,pub,addr = get_wallet(addr_index=addr_index)
    increment()
    resp = {
        "private key": pvk,
        "public key": pub,
        "address": addr
    }
    return jsonify(resp)


@app.route('/api/btc/wallet-balance/', methods=['GET'])
def get_balance():
    addr =  request.json['addr']
    balance = get_wallet_balance(addr)
    return jsonify(balance)
    

if __name__ == '__main__':
    app.run(host ='0.0.0.0',debug=True)