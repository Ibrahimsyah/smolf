import {
  Controller,
  ForbiddenException,
  Get,
  UseFilters,
  UseGuards,
} from '@nestjs/common';
import { HttpExceptionFilter } from 'src/filters/http-exception';
import { AuthGuard } from 'src/guards/auth';

@Controller()
export class HelloController {
  @Get('/hello')
  @UseGuards(AuthGuard)
  getHello(): string {
    return 'heheh';
  }

  @Get('/hi')
  @UseFilters(HttpExceptionFilter)
  hi(): string {
    throw new ForbiddenException();
  }
}
