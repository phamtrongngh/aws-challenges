import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { FileRepository } from './repository/app.repository';

@Module({
  imports: [ConfigModule.forRoot()],
  controllers: [AppController],
  providers: [
    {
      provide: 'Repository',
      useFactory(configService: ConfigService) {
        const filePath = configService.getOrThrow('REPO_FILE_PATH');
        return new FileRepository({ filePath });
      },
      inject: [ConfigService],
    },
  ],
})
export class AppModule {}
