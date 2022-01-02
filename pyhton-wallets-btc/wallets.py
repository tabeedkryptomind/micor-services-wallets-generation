from hdwallet import HDWallet
import requests

PASSPHRASE = "12415135136136"


def generate_address(hdwallet : HDWallet, addr_index: int):
    hdwallet.from_index(44, hardened=True)
    hdwallet.from_index(0, hardened=True)
    hdwallet.from_index(0, hardened=True)
    hdwallet.from_index(0)
    hdwallet.from_index(addr_index)


def get_wallet(addr_index):
    hdwallet = HDWallet(symbol="BTC", use_default_path=False)
    hdwallet.from_mnemonic(
        mnemonic="sceptre capter séquence girafe absolu relatif fleur zoologie muscle sirop saboter parure", 
        passphrase=PASSPHRASE
    )
    generate_address(hdwallet, addr_index)
    return hdwallet.private_key(), hdwallet.public_key(), hdwallet.p2pkh_address()

def get_wallet_balance(addr):
   URL = "https://api.blockcypher.com/v1/btc/main/addrs/{0}/balance".format(addr)
   r = requests.get(url = URL)
   return r.json()