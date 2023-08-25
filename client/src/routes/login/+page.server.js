/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ cookies, request }) => {
		const formData = await request.formData();
		const username = formData.get('username');
		const password = formData.get('password');

		const rawResponse = await fetch('http://localhost:8080/login', {
			method: 'POST',
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({username: username, hash: password})
		});
		if (rawResponse.status !== 200) {
			return {success: false}
		}
		const content = await rawResponse.json();
		cookies.set('session_token', content.session_token);
		return {success: true}
	}
};