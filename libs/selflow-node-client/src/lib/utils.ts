export function mapObject<S, D>(
  obj: Record<string, S>,
  map: (value: S, key: string) => D
) {
  return Object.fromEntries(
    Object.keys(obj).map((key) => [key, map(obj[key], key)])
  );
}
