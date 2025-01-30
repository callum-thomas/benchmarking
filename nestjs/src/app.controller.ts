import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('/simpleHealthcheck')
  getSimpleHealthcheck(): string {
    return 'service healthy';
  }

  @Get('/healthcheck')
  getHealthcheck(): string {
    return JSON.stringify({ status: 'healthy' });
  }
}
