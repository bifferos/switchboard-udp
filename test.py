#!/usr/bin/env python3

import os
import json
import socket
import unittest

CONFIG_PATH = 'config.json'

def load_config():
    with open(CONFIG_PATH) as f:
        return json.load(f)

def udp_query(filename, port):
    server_address = ('localhost', port)
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    try:
        sock.sendto(filename.encode(), server_address)
        data, _ = sock.recvfrom(1024)
        return data.decode()
    finally:
        sock.close()

class UDPFileExistenceTest(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cfg = load_config()
        cls.state_dir = cfg['stateDir']
        cls.port = cfg['port']
        cls.test_file = 'testfile.txt'
        os.makedirs(cls.state_dir, exist_ok=True)
        cls.test_path = os.path.join(cls.state_dir, cls.test_file)

    def test_file_exists(self):
        with open(self.test_path, 'w') as f:
            f.write('test')
        response = udp_query(self.test_file, self.port)
        self.assertEqual(response, 'true')

    def test_file_not_exists(self):
        if os.path.exists(self.test_path):
            os.remove(self.test_path)
        response = udp_query(self.test_file, self.port)
        self.assertEqual(response, 'false')

if __name__ == '__main__':
    unittest.main()
