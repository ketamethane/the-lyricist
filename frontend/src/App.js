import './App.css';
import LyricsPage from './components/LyricsPage'
import SimpleTestComponent from './components/LyricsPage'
// import LiteYouTubeEmbed from 'react-lite-youtube-embed'
// import 'react-lite-youtube-embed/dist/LiteYouTubeEmbed.css'

//import youtube embed and settle dependency issue for furigana component
const App = () => {
  return (
    <div>
      <LyricsPage/>
    </div>
  );
}

export default App


