from http.server import BaseHTTPRequestHandler, HTTPServer

# import socketserver
import ctypes
import threading
import json
import os
import cgi
import sys

import gc

gc_counter=0

path = os.path.join(os.path.abspath("."), "filter.so")
lib = ctypes.CDLL(path)
final_json_data = {}

class Server(BaseHTTPRequestHandler):
    def _set_headers(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        
    def do_HEAD(self):
        self._set_headers()
    
    #returns response to request
    def do_GET(self):
        self._set_headers()
        self.wfile.write(json.dumps(filter_text(message["text"])).encode())
        gc_counter+=1
        if gc_counter%200000 == 0:
        	gc.collect()
      
    #gets request on server
    def do_POST(self):
        ctype, pdict = cgi.parse_header(self.headers.get('content-type'))
        length = int(self.headers.get('content-length'))
        global message
        message = json.loads(self.rfile.read(length))
        self._set_headers()
        self.wfile.write(json.dumps(message).encode())
      

def run(server_class=HTTPServer, handler_class=Server, port=8080):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    
    print ('Server listening on port %d...' % port)
    httpd.serve_forever()


def filter_text(text):
	all_json_data = {}
	merged_json_data = {}

	words = text.split(" ")

	for word in words:
		for i in range(100):
			c_word = ctypes.c_char_p(word.encode('UTF-8'))

			#get topics
			response = lib.get_topics(c_word)
			raw_json_data = ctypes.c_char_p(response).value.decode()	
			json_data = json.loads(raw_json_data)

			#merge topics
			for j in json_data:
				if j not in all_json_data:
					all_json_data[j]=[]
					all_json_data[j].append(json_data[j])
				else:
					if json_data[j] not in all_json_data[j]:
						all_json_data[j].append(json_data[j])

	for i in all_json_data:
		merged_json_data[i] = max(all_json_data[i])
	
	final_json_data['text'] = text
	final_json_data['topics'] = merged_json_data
	
	return str(final_json_data)


def main():
	lib.initialize_go()

	if len(sys.argv) == 2:
		run(port=int(sys.argv[1]))
	else:
		run()
	

if __name__ == "__main__":
	main()