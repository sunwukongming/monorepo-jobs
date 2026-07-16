type AppProps = {
  title: string;
};

export default function App({ title }: AppProps) {
  return (
    <main className="page">
      <h1>{title} Admin</h1>
      <p>后台管理骨架已就绪。API 代理前缀：/api</p>
    </main>
  );
}
