import AlbumItem from "./AlbumItem";

const AlbumList = ({ albums }) => {
  return (
    <section className="flex flex-col m-8 md:m-10">
      <h1 className="text-3xl font-bold">Albums</h1>
      <ul className="grid grid-cols-3 md:grid-cols-6 m-1 md:m-2 p-1 md:p-4">
        {albums.map((album, index) => (
          <AlbumItem key={album.id} album={album} />
        ))}
      </ul>
    </section>
  );
};

export default AlbumList;
