import { Link } from "react-router-dom";
import AlbumItem from "./AlbumItem";

const AlbumOrLogin = ({ albums }) => {
  if (albums) {
    return (
      <ul className="flex flex-col md:grid-cols-6 m-1 md:m-2 p-1 md:p-4">
        {albums.map((album, index) => (
          <AlbumItem key={album.id} album={album} />
        ))}
      </ul>
    );

    return (
      <button>
        <Link to={"/login"}>Login Page</Link>
      </button>
    );
  }
};

const AlbumList = ({ albums }) => {
  return (
    <section className="flex flex-col m-8 md:m-10 font-mono">
      <h1 className="text-3xl font-bold">Albums</h1>
      <AlbumOrLogin albums={albums} />
    </section>
  );
};

export default AlbumList;
