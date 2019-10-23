<?php
declare(strict_types=1);

namespace Hyperf\Support\Traits;

use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Hyperf\Support\Facades\Token;
use Hyperf\Support\Facades\Utils;
use Hyperf\Support\Redis\RefreshToken;
use Hyperf\Utils\Str;

/**
 * Trait Auth
 * @package Hyperf\Support\Traits
 * @property RequestInterface $request
 * @property ResponseInterface $response
 */
trait Auth
{
    /**
     * Set RefreshToken Expires
     * @return int
     */
    protected function __refreshTokenExpires()
    {
        return 604800;
    }

    /**
     * Create Cookie Auth
     * @param string $scene
     * @param array $symbol
     * @return \Psr\Http\Message\ResponseInterface
     */
    protected function __create(string $scene, array $symbol = [])
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
        $tokenString = (string)Token::create($scene, $jti, $ack, $symbol);
        $cookie = Utils::cookie($scene . '_token', $tokenString);
        return $this->response->withCookie($cookie)->json([
            'error' => 0,
            'msg' => 'ok'
        ]);
    }

    /**
     * Auth Verify
     * @param $scene
     * @return array|\Psr\Http\Message\ResponseInterface
     */
    protected function __verify($scene)
    {
        try {
            $tokenString = $this->request->cookie($scene . '_token');
            $result = Token::verify($scene, $tokenString);
            if ($result->expired) {
                /**
                 * @var $token \Lcobucci\JWT\Token
                 */
                $token = $result->token;
                $jti = $token->getClaim('jti');
                $ack = $token->getClaim('ack');
                $verify = RefreshToken::create()->verify($jti, $ack);
                if (!$verify) {
                    return [
                        'error' => 1,
                        'msg' => 'refresh token verification expired'
                    ];
                }
                $symbol = $token->getClaim('symbol');
                $preTokenString = (string)Token::create(
                    $scene . '_token',
                    $jti,
                    $ack,
                    $symbol
                );
                $cookie = Utils::cookie($scene . '_token', $preTokenString);
                return $this->response->withCookie($cookie)->json([
                    'error' => 0,
                    'msg' => 'ok'
                ]);
            }

            return [
                'error' => 0,
                'msg' => 'ok'
            ];
        } catch (\Exception $e) {
            return [
                'error' => 1,
                'msg' => $e->getMessage()
            ];
        }
    }

    /**
     * Destory Auth
     * @param string $scene
     * @return \Psr\Http\Message\ResponseInterface
     */
    protected function __destory(string $scene)
    {
        $tokenString = $this->request->cookie($scene . '_token');
        $token = Token::get($tokenString);
        RefreshToken::create()->clear(
            $token->getClaim('jti'),
            $token->getClaim('ack')
        );
        $cookie = Utils::cookie($scene . '_token', '');
        return $this->response->withCookie($cookie)->json([
            'error' => 0,
            'msg' => 'ok'
        ]);
    }
}