export const baseURL = "http://localhost:8080"

export function getURLString(url: string, base?: string): string {
	if (base) {
		return new URL(url, base).toString()
	}
	return new URL(url, baseURL).toString()
}

export async function GET<T>(url: string, base?: string): Promise<T> {
	const response = await fetch(getURLString(url, base))
	if (!response.ok) {
		console.log(response);

		throw new Error(response.statusText)
	}

	return await (response.json() as Promise<T>)
}

// export async function DELETE<T>(url: string, base?: string): Promise<T> {
// 	const response = await fetch(getURLString(url, base), {
// 		method: "DELETE",
// 	})
// 	if (!response.ok) {
// 		console.log(response);

// 		throw new Error(response.statusText)
// 	}

// 	const res =  await (response.json() as Promise<T>)
// 	return res 
// }

export async function DELETE(url: string, base?: string): Promise<Response>  {
	return await fetch(getURLString(url, base), {
		method: "DELETE",
	});
}