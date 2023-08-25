// @ts-nocheck
/** @type {import('./$types').Actions} */
export const actions = {
	logout: async ({ cookies }) => {
		const rawResponse = await fetch('http://localhost:8080/logout', {
			method: 'POST',
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json',
				'Session-Token': cookies.get('session_token')
			}
		});
		cookies.delete('session_token');
		if (rawResponse.status !== 200) {
			return {}
		}
		return {}
	}
};

/** @type {import('../../../.svelte-kit/types/src/routes').PageServerLoad} */
export async function load({cookies}) {
	const response = await fetch(`http://localhost:8080/user`, {
		method: 'GET',
		headers: {
			'Accept': 'application/json',
			'Content-Type': 'application/json',
			'Session-Token': cookies.get('session_token')
		}
	});
	if (response.status !== 200) {
		return {success: false}
	}
	return {success: true, user: await response.json()};
}
