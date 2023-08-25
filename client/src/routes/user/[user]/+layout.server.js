// @ts-nocheck
/** @type {import('./$types').PageServerLoad} */

export async function load(data) {
	 const response = await fetch(`http://localhost:8080/issue/user/${data.params.user}`);
	 return {issues: await response.json()};
}
