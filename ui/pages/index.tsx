import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";
import Game from "../components/game";
import styles from "../styles/BoardPage.module.css";

const Home: NextPage = () => (
  <div className={styles.main}>
    <Head>
      <title>Free Go Game</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>
    <main>
      <Game gameId={1} />
    </main>
  </div>
);

export default Home;
