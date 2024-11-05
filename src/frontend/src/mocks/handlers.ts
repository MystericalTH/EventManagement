import { http } from 'msw';
import { HttpResponse } from 'msw';
export const handlers = [
	// Mock a GET request to "/api/users"
	http.get('/api/users', ({ params }) => {
		return HttpResponse.json([
			{
				name: 'Nattawin Yamprasert',
				email: 'test@email.com'
			},
			{
				name: 'EIEIE YAYY',
				email: 'test2@email.com'
			}
		]);
	})
];
