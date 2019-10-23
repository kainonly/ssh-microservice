<?php
declare(strict_types=1);

namespace Hyperf\Support\Traits;

use Hyperf\HttpServer\Contract\ResponseInterface;
use Hyperf\Support\Facades\Token;
use Hyperf\Support\Facades\Utils;
use Hyperf\Support\Redis\RefreshToken;
use Hyperf\Utils\Str;

/**
 * Trait Auth
 * @package Hyperf\Support\Traits
 * @property ResponseInterface $response
 */
trait Auth
{
    protected function __auth(string $scene, array $symbol = [])
    {
        $jti = Utils::uuid()->toString();
        $ack = Str::random();
        $result = RefreshToken::create()->factory($jti, $ack, $this->__refreshTokenExpires());
        if (!$result) {
            return $this->response->json([
                'error' => 1,
                'msg' => 'refresh token set failed'
            ]);
        }
        $stringToken = (string)Token::create($scene, $jti, $ack, $symbol);
        $cookie = Utils::cookie($scene . '_token', $stringToken);
        return $this->response->withCookie($cookie)->json([
            'error' => 0,
            'msg' => 'ok'
        ]);
    }

    protected function __refreshTokenExpires()
    {
        return 604800;
    }
}