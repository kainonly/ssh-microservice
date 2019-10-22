<?php
declare(strict_types=1);

namespace Hyperf\Support\Traits;

use Hyperf\HttpServer\Contract\ResponseInterface;
use Hyperf\Support\Facades\Token;

/**
 * Trait Auth
 * @package Hyperf\Support\Traits
 * @property ResponseInterface $response
 */
trait Auth
{
    protected function __auth(string $scene, array $symbol = [])
    {
        $stringToken = (string)Token::create($scene, $symbol);
        return $this->response->json([
            'error' => 0,
            'token' => $stringToken
        ]);
    }
}