import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Login from './Login'; 

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false); 
  const [articles, setArticles] = useState(null); 

  useEffect(() => {
    if (isLoggedIn) {
      async function fetchArticles() {
        try {
          const response = await axios.get('http://localhost:3030/posts');
          if (Array.isArray(response.data.posts)) {
            setArticles(response.data.posts); 
          } else {
            console.error('Response data.posts is not an array:', response.data);
            setArticles([]); 
          }
        } catch (error) {
          console.error('Error fetching articles:', error);
          setArticles([]); 
        }
      }
      fetchArticles();
    }
  }, [isLoggedIn]); 

  const handleLogin = () => {
    setIsLoggedIn(true); 
  };

  if (!isLoggedIn) {
    return <Login onLogin={handleLogin} />;
  }

  return (
    <div className="container">
      <header>
        <h1 className="title">Articles</h1>
      </header>
      <main>
        <ul className="articles-list">
          {articles && articles.map((article) => (
            <li className="article" key={article.id}>
              <div className="article-content">
                <h2 className="article-title">{article.Title}</h2>
                <p className="article-description">{article.Content}</p>
              </div>
              <div className="article-thumbnail">
                <img src={article.Image} alt={article.Title} />
              </div>
            </li>
          ))}
        </ul>
      </main>
      <footer>
        <p className="footer-text">© 2024 Vietnam News</p>
      </footer>
    </div>
  );
}

export default App;
