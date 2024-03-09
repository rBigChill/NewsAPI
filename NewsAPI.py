import apiKey
import subprocess
import sys
import webbrowser

try:
    from newsapi import NewsApiClient
except:
    subprocess.check_call([sys.executable, "-m", "pip", "install", "newsapi-python"])
finally:
    from newsapi import NewsApiClient

class Article:
    def __init__(self):
        self.title = ''
        self.url = ''

class NewsAPI:
    def __init__(self):
        client = NewsApiClient(api_key = apiKey.ApiKey)
        top = client.get_top_headlines(page_size=10,country='us') 
        self.Articles = []

        for i, v in enumerate(top['articles']):
            a = Article()
            a.title = v['title']
            a.url = v['url']
            self.articles.append(a)
            print(f"\n{i+1}) {a.title}\n")

if __name__ == "__main__":
    n = NewsAPI()
    
    selection = 0

    while selection != 'q':
        selection = input("Print article (#) or (q)uit?: ")
        if selection != 'q':
            webbrowser.open(n.Articles[int(selection)-1].url)
