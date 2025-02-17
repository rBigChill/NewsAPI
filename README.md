# News API CLI Reader

Overview
This script fetches and displays the top 10 news articles from the [News API](https://newsapi.org/) for the United States. Users can view article titles in the terminal and open selected articles in their web browser.

Features
- Fetches the latest top 10 US news articles.
- Displays article titles in a formatted list.
- Allows users to open selected articles in their default web browser.
- Automatically installs missing dependencies.

Prerequisites
- Python 3.x
- An API key from [News API](https://newsapi.org/)
- Internet connection

Installation & Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/rBigChill/NewsAPI-CLI.git
   cd NewsAPI-CLI
   ```
2. Install dependencies:
   ```bash
   pip install newsapi-python
   ```
3. Add your API key in the `apiKey.py` file:
   ```python
   ApiKey = "YOUR_NEWSAPI_KEY"
   ```

Usage
Run the script using:
```bash
python newsapi_reader.py
```

#How It Works
- The script fetches top news articles from News API.
- It prints a numbered list of article titles.
- Users can enter an article number to open it in their browser.
- Pressing Enter quits the program.

Future Improvements
- Support for fetching more than 10 articles.
- Additional filtering options (e.g., by category, source, keyword).
- Caching for offline access.

License
This project is open-source and available for modification and distribution.

