#!/usr/bin/env python2
"""
Serveur SPA pour Python 2 (fallback)
"""

import SimpleHTTPServer
import SocketServer
import urlparse
import os
import sys

class SPAHandler(SimpleHTTPServer.SimpleHTTPRequestHandler):
    def do_GET(self):
        url_path = urlparse.urlparse(self.path).path
        file_path = self.translate_path(url_path)
        
        # Fichiers statiques a servir normalement
        static_extensions = ['.js', '.wasm', '.css', '.png', '.jpg', '.jpeg', '.gif', '.svg', '.ico', '.html', '.json', '.txt']
        static_prefixes = ['/wasm_exec.js', '/app.wasm', '/favicon']
        
        # Si c'est un fichier statique, le servir normalement
        is_static_file = (
            any(url_path.endswith(ext) for ext in static_extensions) or
            any(url_path.startswith(prefix) for prefix in static_prefixes) or
            (os.path.exists(file_path) and os.path.isfile(file_path))
        )
        
        if not is_static_file:
            # Pour toutes les routes non-statiques, rediriger vers index.html
            self.path = '/index.html'
        
        return SimpleHTTPServer.SimpleHTTPRequestHandler.do_GET(self)

def main():
    port = int(sys.argv[1]) if len(sys.argv) > 1 else 8080
    
    try:
        httpd = SocketServer.TCPServer(('', port), SPAHandler)
        print('Serveur SPA demarre sur http://localhost:{}'.format(port))
        print('Appuyez sur Ctrl+C pour arreter')
        httpd.serve_forever()
    except KeyboardInterrupt:
        print('\nServeur arrete')
    except Exception as e:
        print('Erreur: {}'.format(e))
        sys.exit(1)

if __name__ == '__main__':
    main()
