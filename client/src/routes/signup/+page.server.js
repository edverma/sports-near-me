/** @type {import('../../../../.svelte-kit/types/src/routes').Actions} */
export const actions = {
	default: async ({ request }) => {
		const formData = await request.formData();
		const firstName = formData.get('first_name');
		const lastName = formData.get('last_name');
		const email = formData.get('email');
		const username = formData.get('username');
		const password = formData.get('password');

		const rawResponse = await fetch('http://localhost:8080/user', {
			method: 'POST',
			headers: {
				'Accept': 'application/json',
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				credential: {
					username: username,
					hash: password
				},
				user: {
					first_name: firstName,
					last_name: lastName,
					email: email,
					username: username
				}
			})});
		if (rawResponse.status !== 200) {
			return {success: false}
		}
		return {success: true}
	}
};