export default function AlbumLikesBy({ users }) {
  return (
    <div className="p-4 m-3 bg-white rounded-lg shadow-lg">
      <h3 className="mb-3 text-xl font-semibold text-gray-900">Liked by</h3>
      <ul className="space-y-2">
        {users.map((user) => (
          <li
            key={user.ID}
            className="px-3 py-2 text-gray-700 bg-gray-100 rounded-lg"
          >
            {user.ID}
          </li>
        ))}
      </ul>
    </div>
  );
}
