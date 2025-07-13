#!/usr/bin/env python3
"""
Serveur SPA (Single Page Application) pour le développement
Redirige toutes les routes non-fichiers vers index.html
"""

import http.server
import socketserver
import urllib.parse
import os
import sys

class SPAHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        # Parse l'URL pour obtenir le chemin
        url_path = urllib.parse.urlparse(self.path).path
        
        # Gestion spéciale pour les assets WebAssembly
        if url_path == '/wasm_exec.js' and os.path.exists('wasm_exec.js'):
            self.path = '/wasm_exec.js'
            return super().do_GET()
        elif url_path == '/app.wasm' and os.path.exists('app.wasm'):
            self.path = '/app.wasm'
            return super().do_GET()
        
        file_path = self.translate_path(url_path)
        
        # Fichiers statiques à servir normalement
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
        
        # Appeler la méthode parent pour servir le fichier
        return super().do_GET()

def main():
    port = int(sys.argv[1]) if len(sys.argv) > 1 else 8080
    
    try:
        with socketserver.TCPServer(('', port), SPAHandler) as httpd:
            print(f'Serveur SPA démarré sur http://localhost:{port}')
            print('Appuyez sur Ctrl+C pour arrêter')
            httpd.serve_forever()
    except KeyboardInterrupt:
        print('\nServeur arrêté')
    except OSError as e:
        if e.errno == 48:  # Address already in use
            print(f'❌ Le port {port} est déjà utilisé')
        else:
            print(f'❌ Erreur: {e}')
        sys.exit(1)

if __name__ == '__main__':
    main()
