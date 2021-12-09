const baseURL = "http://localhost:8080"

export async function GET<T>(url: string, base?: string): Promise<T> {
	let _url: URL
	if(base) {
		_url = new URL(url, base)
	} else {
		_url = new URL(url, baseURL)
	}

	const response = await fetch(_url.toString())
	if (!response.ok) {
		console.log(response);

		throw new Error(response.statusText)
	}
	return await (response.json() as Promise<T>)
}
