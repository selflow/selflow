import {trpc} from "../utils/trpc";

export function Index() {
  const {data} = trpc.status.useQuery("8416d9d7-c328-42c1-b03a-825761800e0d")
  if (!data) {
    return <p>Loading...</p>
  }

  return (
   <pre>{JSON.stringify(data)}</pre>
  );
}

export default Index;
