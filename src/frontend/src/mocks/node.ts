import { setupServer } from 'msw/node';
import { handlers } from './handlers';

// Configure the worker with the defined handlers
export const server = setupServer(...handlers);
