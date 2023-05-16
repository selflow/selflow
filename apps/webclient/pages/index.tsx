import {trpc} from "../utils/trpc";
import {Button} from "@selflow/ui/components-kit";

export function Index() {
  const {data} = trpc.status.useQuery("8416d9d7-c328-42c1-b03a-825761800e0d")
  if (!data) {
    return <p>Loading...</p>
  }

  return (
    <div>
      <pre>{JSON.stringify(data)}</pre>
      <Button/>
    </div>
  );
}

export default Index;
